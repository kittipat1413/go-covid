package server

import (
	"net/http"

	"go-covid/config"
	"go-covid/controllers"
	"go-covid/controllers/render"
	"go-covid/data"
	domainerrors "go-covid/domain/errors"

	"github.com/felixge/httpsnoop"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	geocoderservice "go-covid/services/geocoder"
	patientservice "go-covid/services/patient"

	patientrepo "go-covid/domain/repository/patient"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{cfg}
}

func (s *Server) Start() error {
	router := httprouter.New()

	patientRepo := patientrepo.NewPatientRepository()
	geocoderService := geocoderservice.NewGeocoderService()
	patientService := patientservice.NewPatientService(geocoderService, patientRepo)
	allControllers := []controllers.ControllerInterface{
		controllers.NewPatientController(patientService),
	}

	if err := controllers.MountAll(router, allControllers); err != nil {
		return err
	}

	router.NotFound = http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		render.Error(resp, r, domainerrors.ErrNotFound)
	})

	db, err := data.Connect(s.cfg)
	if err != nil {
		return err
	}

	stack := s.insertCfg(
		s.logRequests(
			s.insertDB(db, router),
		),
	)

	handler := cors.AllowAll().Handler(stack)
	return http.ListenAndServe(s.cfg.ListenAddr(), handler)
}

func (s *Server) logRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		var metrics *httpsnoop.Metrics
		s.cfg.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.RequestURI)
		defer func() {
			if metrics == nil {
				return
			}

			s.cfg.Printf("%s %s %s - HTTP %d %s\n",
				r.RemoteAddr, r.Method, r.RequestURI,
				metrics.Code, metrics.Duration)
		}()

		m := httpsnoop.CaptureMetrics(handler, resp, r)
		metrics = &m
	})
}

func (s *Server) insertCfg(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(resp, config.NewRequest(r, s.cfg))
	})
}

func (s *Server) insertDB(db *sqlx.DB, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(resp, r.WithContext(data.NewContext(r.Context(), db)))
	})
}

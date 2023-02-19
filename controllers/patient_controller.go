package controllers

import (
	"net/http"

	"go-covid/controllers/render"
	"go-covid/model/request"
	"go-covid/model/view"

	patientservice "go-covid/services/patient"

	"github.com/julienschmidt/httprouter"
)

type PatientController struct {
	patientservice patientservice.PatientService
	viewMapper     *view.Mapper
}

func NewPatientController(patientservice patientservice.PatientService) PatientController {
	return PatientController{
		patientservice: patientservice,
		viewMapper:     new(view.Mapper),
	}
}

func (c PatientController) Mount(router *httprouter.Router) error {
	router.POST("/patient/register", c.Register)
	router.GET("/patient/summary", c.Summary)
	return nil
}

func (c PatientController) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	patient := &request.RegisterPatient{}
	if err := ReadJSON(r.Body, patient); err != nil {
		render.Error(w, r, err)
		return
	}
	if err := patient.Validate(); err != nil {
		render.Error(w, r, err)
		return
	}

	registerInput := patientservice.RegisterInput{
		FirstName: patient.FirstName,
		LastName:  patient.LastName,
		Gender:    patient.Gender,
		Age:       patient.Age,
		Lat:       patient.Lat,
		Lng:       patient.Lng,
	}
	if err := c.patientservice.Register(r.Context(), registerInput); err != nil {
		render.Error(w, r, err)
	} else {
		render.JSON(w, r, map[string]string{"message": "your registration is successfully completed"})
	}

}

func (c PatientController) Summary(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	city := r.URL.Query().Get("city")
	if out, err := c.patientservice.Summary(r.Context(), city); err != nil {
		render.Error(w, r, err)
	} else {
		resp := c.viewMapper.ToPatientSummary(out)
		render.JSON(w, r, resp)
	}

}

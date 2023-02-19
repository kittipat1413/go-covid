package heregeocoder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-covid/config"
	"go-covid/domain/entity"
	"net/http"
	"time"
)

var (
	RequestTimeout int = 15
	revgeocodeURL      = "https://revgeocode.search.hereapi.com/v1/revgeocode"
)

type Client struct {
	client *http.Client
	apiKey string
}

func NewClient(ctx context.Context) *Client {
	return &Client{
		client: &http.Client{Timeout: time.Duration(RequestTimeout) * time.Second},
		apiKey: config.FromContext(ctx).HereApiKey(),
	}
}

type Response struct {
	Items []struct {
		Address struct {
			Label       string
			CountryCode string
			CountryName string
			StateCode   string
			State       string
			County      string
			District    string
			City        string
			Street      string
			PostalCode  string
			HouseNumber string
		}
		Position struct {
			Lat float64
			Lng float64
		}
	}
}

func (r *Response) ToLocation() *entity.Location {
	if len(r.Items) == 0 {
		return nil
	}
	res := r.Items[0].Position
	return &entity.Location{
		Lat: res.Lat,
		Lng: res.Lng,
	}
}

func (r *Response) ToAddress() *entity.Address {
	if len(r.Items) == 0 {
		return nil
	}

	res := r.Items[0].Address
	addr := &entity.Address{
		CountryCode: res.CountryCode,
		CountryName: res.CountryName,
		State:       res.State,
		County:      res.County,
		District:    res.District,
		City:        res.City,
		Street:      res.Street,
		PostalCode:  res.PostalCode,
	}
	return addr
}

func (c *Client) GetReverseGeocode(lat, long float64) (response Response, err error) {
	req, err := http.NewRequest(http.MethodGet, revgeocodeURL, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("at", fmt.Sprintf("%f,%f", lat, long))
	query.Add("limit", "1")
	query.Add("lang", "en-US")
	query.Add("apikey", c.apiKey)
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New("error")
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}
	return
}

package entity

import (
	"strings"
)

const (
	Thailand = "THAILAND"
)

type Address struct {
	CountryCode string
	CountryName string
	State       string
	County      string
	District    string
	City        string
	Street      string
	PostalCode  string
}

func (a Address) IsCountryName(countryName string) bool {
	return strings.ToUpper(a.CountryName) == countryName
}

type Location struct {
	Lat float64
	Lng float64
}

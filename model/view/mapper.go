package view

import "go-covid/domain/entity"

type Error struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	HttpCode int         `json:"-"`
}

type PatientSummary struct {
	City         string  `json:"city"`
	District     string  `json:"district,omitempty"`
	Gender       string  `json:"gender"`
	AvgAge       float64 `json:"avg_age"`
	PatientCount uint64  `json:"count"`
}

type Mapper struct{}

func (m Mapper) ToPatientSummary(patientSummary []entity.PatientSummary) (out []PatientSummary) {
	out = make([]PatientSummary, len(patientSummary))
	for i, v := range patientSummary {
		out[i] = PatientSummary{
			City:         v.City,
			District:     v.District,
			Gender:       v.Gender.String(),
			AvgAge:       v.AvgAge,
			PatientCount: v.PatientCount,
		}
	}

	return
}

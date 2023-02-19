package entity

type Patients []Patient
type Patient struct {
	ID        int64
	FirstName string
	LastName  string
	Gender    Gender
	Age       uint8
	City      string
	District  string
	Street    string
	Location
}

type PatientSummary struct {
	City         string
	District     string
	Gender       Gender
	AvgAge       float64
	PatientCount uint64
}

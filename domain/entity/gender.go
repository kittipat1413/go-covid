package entity

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

var genderIndexMapper = map[string]Gender{
	"Male":   Male,
	"Female": Female,
}

var genderStringMapper = map[Gender]string{
	Male:   "Male",
	Female: "Female",
}

func (s Gender) Parse(gender string) Gender {
	return genderIndexMapper[gender]
}

func (s Gender) String() string {
	return genderStringMapper[s]
}

func (s Gender) Is(expected Gender) bool {
	return s.String() == expected.String()
}

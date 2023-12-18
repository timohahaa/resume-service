package resume

type Resume struct {
	Access               ResumeAccessType         `json:"access"`
	Area                 Area                     `json:"area"`
	BirthDate            string                   `json:"birth_date"`
	BusinessTripReadines BusinessTripReadinesType `json:"business_trip_readiness"`
	Citizenship          []Citizenship            `json:"citizenship"`
	Contact              []Contact                `json:"contact"`
	//    Education <- TODO
	Employments []EmploymentType `json:"employments"`
	Experience  []ExperienceType `json:"experience"`
}

type ResumeAccessType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Area struct {
	Id string `json:"id"`
}

type BusinessTripReadinesType struct {
	Id string `json:"id"`
}

type Citizenship struct {
	Id string `json:"id"`
}

type ContactType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Email string

type PhoneNumber struct {
	City      string `json:"city"`
	Country   string `json:"country"`
	Formatted string `json:"formatted"`
	Number    string `json:"number"`
}

type Contact struct {
	Type  ContactType `json:"type"`
	Value interface{} `json:"value"`
}

type EmploymentType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ExperienceType struct {
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	Description string `json:"description"`
	Position    string `json:"position"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

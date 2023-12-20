package resume

type Resume struct {
	Access               ResumeAccessType         `json:"access"`
	Area                 Area                     `json:"area"`
	BirthDate            string                   `json:"birth_date"`
	BusinessTripReadines BusinessTripReadinesType `json:"business_trip_readiness"`
	Citizenship          []Citizenship            `json:"citizenship"`
	Contacts             []Contact                `json:"contact"`
	Education            Education                `json:"education"`
	Employments          []EmploymentType         `json:"employments"`
	Experience           []ExperienceType         `json:"experience"`
	FirstName            string                   `json:"first_name"`
	Gender               Gender                   `json:"gender"`
	Language             Language                 `json:"language"`
	LastName             string                   `json:"last_name"`
	ProfessionalRoles    []ProfessionalRole       `json:"professional_roles"`
	ResumeLocale         ResumeLocaleType         `json:"resume_locale"`
	ResumeTitle          string                   `json:"title"`
	Schedules            []Schedule               `json:"schedules"`
	Site                 Site                     `json:"site"`
	SkillSet             []string                 `json:"skill_set"`
	Skills               string                   `json:"skills"`
	TravelTime           TravelTimeType           `json:"travel_time"`
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

type Gender struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type LanguageLevel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Language struct {
	Id    string        `json:"id"`
	Name  string        `json:"name"`
	Level LanguageLevel `json:"level"`
}

type ProfessionalRole struct {
	Id string `json:"id"`
}

type ResumeLocaleType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TravelTimeType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SiteType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Site struct {
	Type SiteType `json:"type"`
	URL  string   `json:"url"`
}

type EducationLevelType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PrimaryEducation struct {
	UniName        string `json:"name"`
	GraduationYear string `json:"year"`
	Faculty        string `json:"organization"`
	Specialty      string `json:"result"`
}

type Education struct {
	Level   EducationLevelType `json:"level"`
	Primary []PrimaryEducation `json:"primary"`
}

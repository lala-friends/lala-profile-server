package domain

type PersonProductProject struct {
	PersonId  int    `json:"personId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Introduce string `json:"introduce"`
	ImageUrl  string `json:"imageUrl"`
	RepColor  string `json:"repColor"`
	Blog      string `json:"blog"`
	Github    string `json:"github"`
	Facebook  string `json:"facebook"`

	Projects []Project `json:"projects"`

	Products [] Product `json:"products"`
}

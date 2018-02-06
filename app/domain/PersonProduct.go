package domain

type PersonProduct struct {
	PersonId  int    `json:"personId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Introduce string `json:"introduce"`
	ImageUrl  string `json:"imageUrl"`
	RepColor  string `json:"repColor"`
	Blog      string `json:"blog"`
	Github    string `json:"github"`
	Facebook  string `json:"facebook"`

	Products []Product `json:"products"`
}

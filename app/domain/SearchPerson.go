package domain

type SearchPerson struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Introduce string `json:"introduce"`
	ImageUrl  string `json:"imageUrl"`
	RepColor  string `json:"repColor"`
	Blog      string `json:"blog"`
	Github    string `json:"github"`
	Facebook  string `json:"facebook"`
	Tags      string `json:"tags"`
}

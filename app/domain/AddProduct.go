package domain

type AddProduct struct {
	Name       string   `json:"name"`
	Introduce  string   `json:"introduce"`
	Techs      []string `json:"techs"`
	RepColor   string   `json:"repColor"`
	ImageUrl   string   `json:"imageUrl"`
	Developers [] int   `json:"developers"`
	Details    []Detail `json:"details"`
}

type Detail struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	ProductId   int    `json:"productId"`
}

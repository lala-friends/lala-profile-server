package domain

type ProductPerson struct {
	ProductId int      `json:"productId"`
	Name      string   `json:"name"`
	Introduce string   `json:"introduce"`
	Tech      []string `json:"tech"`
	ImageUrl  string   `json:"imageUrl"`

	Persons	  []Person `json:"persons"`
}


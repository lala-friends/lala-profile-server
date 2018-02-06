package domain

type Product struct {
	ProductId int      `json:"productId"`
	Name      string   `json:"name"`
	Introduce string   `json:"introduce"`
	Tech      []string `json:"tech"`
	ImageUrl  string   `json:"imageUrl"`
}

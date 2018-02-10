package domain

type ProductProductDetailComment struct {
	ProductId int      `json:"productId"`
	Name      string   `json:"name"`
	Introduce string   `json:"introduce"`
	Tech      []string `json:"techs"`
	RepColor  string   `json:"repColor"`
	ImageUrl  string   `json:"imageUrl"`

	ProductDetails []ProductDetail `json:"productDetails"`

	Comments []Comment `json:"comments"`
}
package domain

type ProductDetail struct {
	ProductDetailId int    `json:"productDetailId"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	ImageUrl        string `json:"imageUrl"`
	ProductId       int    `json:"productId"`
}
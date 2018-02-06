package domain

import "database/sql"

type Product struct {
	ProductId int            `json:"productId"`
	Name      string         `json:"name"`
	Introduce sql.NullString `json:"introduce"`
	Tech      []string       `json:"tech"`
	ImageUrl  sql.NullString `json:"imageUrl"`
}

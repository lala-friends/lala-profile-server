package domain

import "database/sql"

type PersonProduct struct {
	PersonId  int            `json:"personId"`
	Name      string         `json:"name"`
	Email     sql.NullString `json:"email"`
	Introduce sql.NullString `json:"introduce"`
	ImageUrl  sql.NullString `json:"imageUrl"`
	RepColor  sql.NullString `json:"repColor"`
	Blog      sql.NullString `json:"blog"`
	Github    sql.NullString `json:"github"`
	Facebook  sql.NullString `json:"facebook"`

	Products []Product `json:"products"`
}

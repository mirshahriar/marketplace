// Package types holds all the types includes API, database, contracts
package types

import "time"

type ProductBody struct {
	Name        string  `json:"name" structs:"name"`
	Description string  `json:"description" structs:"description"`
	Price       float64 `json:"price" structs:"price"`
}

type ProductResponse struct {
	ID uint `json:"id"`
	ProductBody
	CreatedAt time.Time
	UpdatedAt time.Time
}

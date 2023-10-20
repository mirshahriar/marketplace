// Package types holds all the types includes API, database, contracts
package types

import "time"

// ProductBody is the API request body for product
type ProductBody struct {
	// Name is the name of the product. Max length is 100.
	// +required
	Name string `json:"name" structs:"name"`
	// Description is the description of the product. Max length is 200.
	// +required
	Description string `json:"description" structs:"description"`
	// Price is the price of the product. Must be positive number.
	// +required
	Price float64 `json:"price" structs:"price"`
}

// ProductResponse is the API response for product
type ProductResponse struct {
	// ID represents the product id
	ID uint `json:"id"`
	// ProductBody is the product body (name, description, price)
	ProductBody
	// CreatedAt is the time when the product was created
	CreatedAt time.Time
	// UpdatedAt is the time when the product was updated
	UpdatedAt time.Time
}

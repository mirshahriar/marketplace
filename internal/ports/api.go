// Package ports declares all PORT interface
package ports

import (
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// APIPort is the interface for all API ports
type APIPort interface {
	// ListProduct returns a list of products with pagination and sorting
	ListProduct(page types.PageReq, sort types.SortReq) (types.Page[types.ProductResponse], errors.Error)
	// CreateProduct creates a new product and returns the created product
	CreateProduct(body types.ProductBody) (types.ProductResponse, errors.Error)
	// GetProduct returns a product by its id
	GetProduct(productID uint) (types.ProductResponse, errors.Error)
	// UpdateProduct updates a product
	UpdateProduct(productID uint, body types.ProductBody) errors.Error
	// DeleteProduct deletes a product
	DeleteProduct(productID uint) errors.Error

	// GetUserByToken returns a user by its token. This is used by authentication middleware
	GetUserByToken(token string) (types.LoggedInUser, bool, errors.Error)
}

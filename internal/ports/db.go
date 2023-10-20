package ports

import (
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// DBPort is the interface for all DB ports
type DBPort interface {
	// ListProduct returns a list of products with pagination and sorting
	ListProduct(page types.PageReq, sort types.SortReq) ([]types.Product, errors.Error)
	// InsertProduct inserts a new product and returns the created product
	InsertProduct(product types.Product) (types.Product, errors.Error)
	// GetProductByID returns a product by its id
	GetProductByID(productID uint) (types.Product, errors.Error)
	// UpdateProduct updates a product with the given update map
	UpdateProduct(productID uint, update map[string]interface{}) errors.Error
	// DeleteProduct deletes a product
	DeleteProduct(productID uint) errors.Error
	// CountProduct returns the total number of products which is used for pagination
	CountProduct() (int, errors.Error)

	// GetUserByToken returns a user by its token. This is used by authentication middleware
	GetUserByToken(token string) (types.User, bool, errors.Error)
}

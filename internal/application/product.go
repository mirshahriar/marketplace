package app

import (
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// CreateProduct creates a new product and returns the created product
func (a Adapter) CreateProduct(body types.ProductBody) (types.ProductResponse, errors.Error) {
	// InsertProduct inserts a new product into the database
	product, cErr := a.db.InsertProduct(body.ToProduct())
	if cErr != nil {
		return types.ProductResponse{}, cErr
	}

	return product.ToProductResponse(), nil
}

// ListProduct returns a list of products with pagination and sorting
func (a Adapter) ListProduct(page types.PageReq, sort types.SortReq) (types.Page[types.ProductResponse], errors.Error) {
	// ListProduct returns a list of products with pagination and sorting from the database
	products, cErr := a.db.ListProduct(page, sort)
	if cErr != nil {
		return types.Page[types.ProductResponse]{}, cErr
	}

	var resp []types.ProductResponse
	for _, product := range products {
		resp = append(resp, product.ToProductResponse())
	}

	// CountProduct returns the total number of products which is used for pagination
	total, err := a.db.CountProduct()
	if err != nil {
		return types.Page[types.ProductResponse]{}, err
	}

	return types.Page[types.ProductResponse]{
		Data:  resp,
		Page:  page.Page,
		Size:  len(resp),
		Total: total,
	}, nil
}

// GetProduct returns a product by its id
func (a Adapter) GetProduct(productID uint) (types.ProductResponse, errors.Error) {
	product, cErr := a.db.GetProductByID(productID)
	if cErr != nil {
		return types.ProductResponse{}, cErr
	}

	return product.ToProductResponse(), nil
}

// UpdateProduct updates a product by its ID
func (a Adapter) UpdateProduct(productID uint, body types.ProductBody) errors.Error {
	// First we check if the product exists
	_, cErr := a.db.GetProductByID(productID)
	if cErr != nil {
		return cErr
	}

	// UpdateProduct updates a product with map[string]interface{}
	cErr = a.db.UpdateProduct(productID, body.ToMap())
	if cErr != nil {
		return cErr
	}

	return nil
}

// DeleteProduct deletes a product by its ID
func (a Adapter) DeleteProduct(productID uint) errors.Error {
	// First we check if the product exists
	_, cErr := a.db.GetProductByID(productID)
	if cErr != nil {
		return cErr
	}

	if cErr = a.db.DeleteProduct(productID); cErr != nil {
		return cErr
	}

	return nil
}

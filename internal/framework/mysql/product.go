package db

import (
	gError "errors"

	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
	"gorm.io/gorm"
)

// InsertProduct inserts a new product into the database
func (a Adapter) InsertProduct(product types.Product) (types.Product, errors.Error) {
	err := a.db.Create(&product).Error
	if err != nil {
		return types.Product{}, errors.InternalDBError(err)
	}

	return product, nil
}

// ListProduct returns a list of products with pagination and sorting
func (a Adapter) ListProduct(page types.PageReq, sort types.SortReq) ([]types.Product, errors.Error) {
	var products []types.Product

	db := a.db.Model(&types.Product{})
	db = db.Scopes(page.Paginate(), sort.Sort())

	if err := db.Find(&products).Error; err != nil {
		return nil, errors.InternalDBError(err)
	}

	return products, nil
}

// GetProductByID returns a product by its ID
func (a Adapter) GetProductByID(productID uint) (types.Product, errors.Error) {
	var product types.Product

	if err := a.db.First(&product, productID).Error; err != nil {
		// If the product is not found, we return a NoEntityError
		if gError.Is(err, gorm.ErrRecordNotFound) {
			return types.Product{}, errors.NoEntityError("product")
		}

		return types.Product{}, errors.InternalDBError(err)
	}

	return product, nil
}

// UpdateProduct updates a product by its ID
func (a Adapter) UpdateProduct(productID uint, update map[string]interface{}) errors.Error {
	var productModel types.Product

	if err := a.db.Model(&productModel).Where("id = ?", productID).Updates(update).Error; err != nil {
		return errors.InternalDBError(err)
	}

	return nil
}

// DeleteProduct deletes a product by its ID
func (a Adapter) DeleteProduct(productID uint) errors.Error {
	var productModel types.Product

	if err := a.db.Delete(&productModel, productID).Error; err != nil {
		return errors.InternalDBError(err)
	}

	return nil
}

// CountProduct returns the number of products in the database
func (a Adapter) CountProduct() (int, errors.Error) {
	var count int64

	if err := a.db.Model(&types.Product{}).Count(&count).Error; err != nil {
		return 0, errors.InternalDBError(err)
	}

	return int(count), nil
}

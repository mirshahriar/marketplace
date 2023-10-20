package types

import "github.com/fatih/structs"

// ToProduct converts ProductBody (request) to Product (model)
func (p ProductBody) ToProduct() Product {
	return Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

// ToProductResponse converts Product (model) to ProductResponse (response)
func (p Product) ToProductResponse() ProductResponse {
	return ProductResponse{
		ID: p.ID,
		ProductBody: ProductBody{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToMap converts ProductBody (request) to map[string]interface{} (used for gorm update)
func (p ProductBody) ToMap() map[string]interface{} {
	return structs.Map(p)
}

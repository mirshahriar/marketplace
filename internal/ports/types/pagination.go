package types

import (
	"fmt"

	"gorm.io/gorm"
)

// Page is used to return paginated data
type Page[T any] struct {
	// Data holds the original response list
	Data []T `json:"data"`
	// Sort represents the applied sorting
	Sort SortReq `json:"sort"`
	// Page represents the current page
	Page int `json:"page,omitempty"`
	// Size represents the number of items in a page
	Size int `json:"size,omitempty"`
	// Total represents the total number of items in the DB
	Total int `json:"total,omitempty"`
}

// PageReq is used to paginate data
type PageReq struct {
	// Page represents the requested page
	Page int `json:"page" query:"page"`
	// Size represents the number of items in a page
	Size int `json:"size" query:"size"`
}

// Paginate returns scope for pagination
func (p *PageReq) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((p.Page - 1) * p.Size).Limit(p.Size)
	}
}

// NewPageReq returns a new PageReq with default values
func NewPageReq(size int) PageReq {
	return PageReq{1, size}
}

// SortReq is used to sort data
type SortReq struct {
	// By represents the field to sort by
	By string `json:"sort_by" query:"sort_by"`
	// Direction represents the direction of sorting
	Direction SortDirection `json:"sort_direction" query:"sort_direction"`
}

// SortDirection represents the direction of sorting
type SortDirection string

const (
	// ASC represents ascending order
	ASC SortDirection = "ASC"
	// DESC represents descending order
	DESC SortDirection = "DESC"
)

// Sort returns scope for sorting
func (s *SortReq) Sort() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", s.By, s.Direction))
	}
}

// NewSortReq returns a new SortReq with default values
func NewSortReq() SortReq {
	return SortReq{"id", ASC}
}

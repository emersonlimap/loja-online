package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Category    string         `json:"category" gorm:"not null"` // Camiseta, Calça, Vestido, etc.
	Brand       string         `json:"brand"`
	Price       float64        `json:"price" gorm:"not null"`
	CostPrice   float64        `json:"cost_price"`
	SKU         string         `json:"sku" gorm:"unique;not null"`
	Color       string         `json:"color"`
	Size        string         `json:"size"`
	Material    string         `json:"material"`
	Gender      string         `json:"gender"` // Masculino, Feminino, Unissex
	Season      string         `json:"season"` // Verão, Inverno, etc.
	Active      bool           `json:"active" gorm:"default:true"`
	ImageURL    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relacionamentos
	InventoryItems []InventoryItem `json:"inventory_items,omitempty"`
	SaleItems      []SaleItem      `json:"sale_items,omitempty"`
}

type ProductCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Category    string  `json:"category" binding:"required"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CostPrice   float64 `json:"cost_price"`
	SKU         string  `json:"sku" binding:"required"`
	Color       string  `json:"color"`
	Size        string  `json:"size"`
	Material    string  `json:"material"`
	Gender      string  `json:"gender"`
	Season      string  `json:"season"`
	ImageURL    string  `json:"image_url"`
}

type ProductUpdate struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Category    *string  `json:"category"`
	Brand       *string  `json:"brand"`
	Price       *float64 `json:"price"`
	CostPrice   *float64 `json:"cost_price"`
	Color       *string  `json:"color"`
	Size        *string  `json:"size"`
	Material    *string  `json:"material"`
	Gender      *string  `json:"gender"`
	Season      *string  `json:"season"`
	ImageURL    *string  `json:"image_url"`
	Active      *bool    `json:"active"`
}

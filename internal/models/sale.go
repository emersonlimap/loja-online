package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CustomerID    uint           `json:"customer_id"`
	UserID        uint           `json:"user_id" gorm:"not null"` // Vendedor
	TotalAmount   float64        `json:"total_amount" gorm:"not null"`
	Discount      float64        `json:"discount" gorm:"default:0"`
	FinalAmount   float64        `json:"final_amount" gorm:"not null"`
	Status        string         `json:"status" gorm:"default:'pending'"` // pending, confirmed, shipped, delivered, cancelled
	PaymentMethod string         `json:"payment_method"`                  // cash, card, pix, etc.
	Notes         string         `json:"notes"`
	SaleDate      time.Time      `json:"sale_date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`

	// Relacionamentos
	Customer  Customer   `json:"customer"`
	User      User       `json:"user"`
	SaleItems []SaleItem `json:"sale_items"`
}

type SaleItem struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	SaleID     uint    `json:"sale_id" gorm:"not null"`
	ProductID  uint    `json:"product_id" gorm:"not null"`
	Quantity   int     `json:"quantity" gorm:"not null"`
	UnitPrice  float64 `json:"unit_price" gorm:"not null"`
	TotalPrice float64 `json:"total_price" gorm:"not null"`

	// Relacionamentos
	Sale    Sale    `json:"sale,omitempty"`
	Product Product `json:"product"`
}

type SaleCreate struct {
	CustomerID    uint             `json:"customer_id"`
	PaymentMethod string           `json:"payment_method" binding:"required"`
	Discount      float64          `json:"discount"`
	Notes         string           `json:"notes"`
	Items         []SaleItemCreate `json:"items" binding:"required,dive"`
}

type SaleItemCreate struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}

type SaleUpdate struct {
	Status        *string `json:"status"`
	PaymentMethod *string `json:"payment_method"`
	Notes         *string `json:"notes"`
}

type SalesReport struct {
	Period      string  `json:"period"`
	TotalSales  int     `json:"total_sales"`
	TotalAmount float64 `json:"total_amount"`
	TopProducts []struct {
		ProductName string  `json:"product_name"`
		Quantity    int     `json:"quantity"`
		Revenue     float64 `json:"revenue"`
	} `json:"top_products"`
	SalesByStatus map[string]int `json:"sales_by_status"`
	SalesByMonth  []struct {
		Month  string  `json:"month"`
		Count  int     `json:"count"`
		Amount float64 `json:"amount"`
	} `json:"sales_by_month"`
}

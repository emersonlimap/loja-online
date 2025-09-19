package models

import (
	"time"

	"gorm.io/gorm"
)

type InventoryItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null;default:0"`
	MinStock  int            `json:"min_stock" gorm:"default:0"`
	MaxStock  int            `json:"max_stock" gorm:"default:1000"`
	Location  string         `json:"location"` // Localização no estoque
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relacionamentos
	Product         Product             `json:"product"`
	MovementHistory []InventoryMovement `json:"movement_history,omitempty" gorm:"foreignKey:ProductID"`
}

type InventoryMovement struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	ProductID     uint      `json:"product_id" gorm:"not null"`
	MovementType  string    `json:"movement_type" gorm:"not null"` // entry, exit, adjustment
	Quantity      int       `json:"quantity" gorm:"not null"`
	PreviousStock int       `json:"previous_stock"`
	NewStock      int       `json:"new_stock"`
	Reason        string    `json:"reason"`
	UserID        uint      `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`

	// Relacionamentos
	Product Product `json:"product"`
	User    User    `json:"user"`
}

type InventoryUpdate struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
	Reason    string `json:"reason" binding:"required"`
}

type InventoryAdjustment struct {
	ProductID   uint   `json:"product_id" binding:"required"`
	NewQuantity int    `json:"new_quantity" binding:"required,gte=0"`
	Reason      string `json:"reason" binding:"required"`
}

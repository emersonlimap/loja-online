package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique"`
	Phone     string         `json:"phone"`
	CPF       string         `json:"cpf" gorm:"unique"`
	Gender    string         `json:"gender"`
	BirthDate *time.Time     `json:"birth_date"`
	Active    bool           `json:"active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Endereços
	Addresses []Address `json:"addresses,omitempty"`

	// Relacionamentos
	Sales []Sale `json:"sales,omitempty"`
}

type Address struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	CustomerID   uint   `json:"customer_id"`
	Street       string `json:"street" gorm:"not null"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city" gorm:"not null"`
	State        string `json:"state" gorm:"not null"`
	ZipCode      string `json:"zip_code" gorm:"not null"`
	IsDefault    bool   `json:"is_default" gorm:"default:false"`
}

type CustomerCreate struct {
	Name      string     `json:"name" binding:"required"`
	Email     string     `json:"email" binding:"omitempty,email"`
	Phone     string     `json:"phone"`
	CPF       string     `json:"cpf" binding:"required"`
	Gender    string     `json:"gender"`
	BirthDate *time.Time `json:"birth_date"`
}

type CustomerUpdate struct {
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Phone     *string    `json:"phone"`
	Gender    *string    `json:"gender"`
	BirthDate *time.Time `json:"birth_date"`
	Active    *bool      `json:"active"`
}

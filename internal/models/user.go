package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	Role      string         `json:"role" gorm:"default:'user'"` // admin, manager, user
	Active    bool           `json:"active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Permissões por módulo
	Permissions UserPermissions `json:"permissions" gorm:"embedded"`
}

type UserPermissions struct {
	Products  bool `json:"products" gorm:"default:false"`  // Cadastro de produtos
	Customers bool `json:"customers" gorm:"default:false"` // Cadastro de clientes
	Inventory bool `json:"inventory" gorm:"default:false"` // Controle de estoque
	Sales     bool `json:"sales" gorm:"default:false"`     // Controle de vendas
	Reports   bool `json:"reports" gorm:"default:false"`   // Relatórios
	Users     bool `json:"users" gorm:"default:false"`     // Controle de usuários
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

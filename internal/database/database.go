package database

import (
	"log"
	"loja-online/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect estabelece conexão com o banco de dados
func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Migrate executa as migrações do banco de dados
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Customer{},
		&models.Address{},
		&models.InventoryItem{},
		&models.InventoryMovement{},
		&models.Sale{},
		&models.SaleItem{},
	)
}

// CreateDefaultAdmin cria o usuário admin padrão se não existir
func CreateDefaultAdmin(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.User{}).Where("email = ?", "leozinsurfwear@gmail.com").Count(&count).Error; err != nil {
		return err
	}

	// Se usuário já existe, não faz nada
	if count > 0 {
		log.Println("Usuário admin já existe")
		return nil
	}

	// Gera hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("leozin@123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Cria usuário admin
	admin := models.User{
		Name:     "Administrador LeoZin",
		Email:    "leozinsurfwear@gmail.com",
		Password: string(hashedPassword),
		Role:     "admin",
		Active:   true,
		Permissions: models.UserPermissions{
			Products:  true,
			Customers: true,
			Inventory: true,
			Sales:     true,
			Reports:   true,
			Users:     true,
		},
	}

	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	log.Println("✅ Usuário admin criado com sucesso!")
	log.Println("📧 Email: leozinsurfwear@gmail.com")
	log.Println("🔐 Senha: leozin@123")

	return nil
}

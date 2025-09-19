package handlers

import (
	"loja-online/internal/config"

	"gorm.io/gorm"
)

// Handler struct que contém as dependências para os handlers
type Handler struct {
	DB     *gorm.DB
	Config *config.Config
}

// New cria uma nova instância do Handler
func New(db *gorm.DB, config *config.Config) *Handler {
	return &Handler{
		DB:     db,
		Config: config,
	}
}

package main

import (
	"log"
	"os"

	"loja-online/internal/api"
	"loja-online/internal/config"
	"loja-online/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Loja Online API
// @version 1.0
// @description API para sistema de loja online de roupas
// @termsOfService http://swagger.io/terms/

// @contact.name Suporte API
// @contact.email suporte@loja-online.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite 'Bearer ' seguido do token JWT

func main() {
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Printf("Aviso: Arquivo .env não encontrado (%v), usando variáveis do sistema", err)
	} else {
		log.Println("✓ Arquivo .env carregado com sucesso")
	}

	// Inicializa configuração
	cfg := config.Load()

	// Conecta ao banco de dados
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados:", err)
	}

	// Executa migrações
	if err := database.Migrate(db); err != nil {
		log.Fatal("Falha ao executar migrações:", err)
	}

	// Cria usuário admin padrão se não existir
	if err := database.CreateDefaultAdmin(db); err != nil {
		log.Printf("Aviso: Falha ao criar usuário admin padrão: %v", err)
	}

	// Configura Gin
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializa router
	router := api.SetupRouter(db, cfg)

	// Inicia servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta %s", port)
	log.Fatal(router.Run(":" + port))
}

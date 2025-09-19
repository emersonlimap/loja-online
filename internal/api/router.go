package api

import (
	"loja-online/internal/config"
	"loja-online/internal/handlers"
	"loja-online/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter configura todas as rotas da aplicação
func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// Middleware global
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// Servir arquivos estáticos
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")

	// Handlers
	h := handlers.New(db, cfg)

	// Rotas públicas
	authPublic := router.Group("/api/v1/auth")
	{
		authPublic.POST("/login", h.Login)
		authPublic.POST("/register", h.Register)
	}

	// Rotas protegidas
	api := router.Group("/api/v1")
	api.Use(middleware.AuthRequired(cfg.JWTSecret))
	{
		// Auth protegido
		authProtected := api.Group("/auth")
		{
			authProtected.GET("/me", h.Me)
		}
		// Produtos
		products := api.Group("/products")
		{
			products.GET("", h.GetProducts)
			products.POST("", h.CreateProduct)
			products.GET("/:id", h.GetProduct)
			products.PUT("/:id", h.UpdateProduct)
			products.DELETE("/:id", h.DeleteProduct)
		}

		// Clientes
		customers := api.Group("/customers")
		{
			customers.GET("", h.GetCustomers)
			customers.POST("", h.CreateCustomer)
			customers.GET("/:id", h.GetCustomer)
			customers.PUT("/:id", h.UpdateCustomer)
			customers.DELETE("/:id", h.DeleteCustomer)
		}

		// Vendas
		sales := api.Group("/sales")
		{
			sales.GET("", h.GetSales)
			sales.POST("", h.CreateSale)
			sales.GET("/:id", h.GetSale)
			sales.PUT("/:id", h.UpdateSale)
		}

		// Estoque
		inventory := api.Group("/inventory")
		{
			inventory.GET("", h.GetInventory)
			inventory.POST("/adjust", h.AdjustInventory)
			inventory.GET("/movements/:product_id", h.GetInventoryMovements)
		}

		// Relatórios
		reports := api.Group("/reports")
		{
			reports.GET("/sales", h.GetSalesReport)
		}

		// Usuários
		users := api.Group("/users")
		{
			users.GET("", h.GetUsers)
			users.POST("", h.CreateUser)
			users.PUT("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}
	}

	// Rotas web públicas
	router.GET("/", h.Dashboard)
	router.GET("/login", h.LoginPage)
	router.GET("/dashboard", h.Dashboard)
	router.GET("/dashboard", h.Dashboard)

	return router
}

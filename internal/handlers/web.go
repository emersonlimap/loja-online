package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dashboard renderiza a página principal do dashboard
func (h *Handler) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Dashboard - Loja Online",
	})
}

// LoginPage renderiza a página de login
func (h *Handler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login - Loja Online",
	})
}

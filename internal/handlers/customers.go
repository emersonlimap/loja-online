package handlers

import (
	"net/http"
	"strconv"

	"loja-online/internal/models"

	"github.com/gin-gonic/gin"
)

// GetCustomers retorna todos os clientes
func (h *Handler) GetCustomers(c *gin.Context) {
	var customers []models.Customer

	if err := h.DB.Preload("Addresses").Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// GetCustomer retorna um cliente específico
func (h *Handler) GetCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var customer models.Customer
	if err := h.DB.Preload("Addresses").First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// CreateCustomer cria um novo cliente
func (h *Handler) CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"customer": customer})
}

// UpdateCustomer atualiza um cliente existente
func (h *Handler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var customer models.Customer
	if err := h.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	var updateData models.Customer
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Model(&customer).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// DeleteCustomer deleta um cliente
func (h *Handler) DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.DB.Delete(&models.Customer{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente deletado com sucesso"})
}

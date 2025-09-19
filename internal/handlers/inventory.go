package handlers

import (
	"net/http"
	"strconv"

	"loja-online/internal/models"

	"github.com/gin-gonic/gin"
)

// GetInventory retorna todos os itens do inventário
func (h *Handler) GetInventory(c *gin.Context) {
	var inventoryItems []models.InventoryItem

	if err := h.DB.Preload("Product").Find(&inventoryItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar inventário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inventory": inventoryItems})
}

// AdjustInventory ajusta a quantidade de um produto no estoque
func (h *Handler) AdjustInventory(c *gin.Context) {
	var adjustment models.InventoryAdjustment

	if err := c.ShouldBindJSON(&adjustment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Busca o item no inventário
	var inventoryItem models.InventoryItem
	if err := h.DB.Where("product_id = ?", adjustment.ProductID).First(&inventoryItem).Error; err != nil {
		// Se não existir, cria um novo
		inventoryItem = models.InventoryItem{
			ProductID: adjustment.ProductID,
			Quantity:  0,
		}
		if err := h.DB.Create(&inventoryItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar item no inventário"})
			return
		}
	}

	// Salva quantidade anterior
	previousQuantity := inventoryItem.Quantity

	// Atualiza quantidade
	inventoryItem.Quantity = adjustment.NewQuantity

	// Salva alterações
	if err := h.DB.Save(&inventoryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ajustar inventário"})
		return
	}

	// Registra movimento
	movementType := "adjustment"
	if adjustment.NewQuantity > previousQuantity {
		movementType = "entry"
	} else if adjustment.NewQuantity < previousQuantity {
		movementType = "exit"
	}

	movement := models.InventoryMovement{
		ProductID:     adjustment.ProductID,
		MovementType:  movementType,
		Quantity:      adjustment.NewQuantity - previousQuantity,
		PreviousStock: previousQuantity,
		NewStock:      adjustment.NewQuantity,
		Reason:        adjustment.Reason,
	}

	// Adiciona usuário se disponível no contexto
	if userID, exists := c.Get("user_id"); exists {
		if uid, ok := userID.(float64); ok {
			movement.UserID = uint(uid)
		}
	}

	if err := h.DB.Create(&movement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar movimento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inventário ajustado com sucesso",
		"item":    inventoryItem,
	})
}

// GetInventoryMovements retorna o histórico de movimentos de um produto
func (h *Handler) GetInventoryMovements(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto inválido"})
		return
	}

	var movements []models.InventoryMovement
	if err := h.DB.Where("product_id = ?", productID).
		Preload("Product").
		Preload("User").
		Order("created_at DESC").
		Find(&movements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar movimentos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movements": movements})
}

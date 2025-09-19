package handlers

import (
	"net/http"
	"strconv"
	"time"

	"loja-online/internal/models"

	"github.com/gin-gonic/gin"
)

// GetSales retorna todas as vendas
func (h *Handler) GetSales(c *gin.Context) {
	var sales []models.Sale

	if err := h.DB.Preload("Customer").Preload("User").Preload("SaleItems.Product").Find(&sales).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar vendas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sales": sales})
}

// GetSale retorna uma venda específica
func (h *Handler) GetSale(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var sale models.Sale
	if err := h.DB.Preload("Customer").Preload("User").Preload("SaleItems.Product").First(&sale, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Venda não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sale": sale})
}

// CreateSale cria uma nova venda
func (h *Handler) CreateSale(c *gin.Context) {
	var sale models.Sale

	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define a data da venda
	sale.SaleDate = time.Now()

	// Calcula o valor final
	sale.FinalAmount = sale.TotalAmount - sale.Discount

	// Inicia transação
	tx := h.DB.Begin()

	if err := tx.Create(&sale).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar venda"})
		return
	}

	// Atualiza o estoque para cada item vendido
	for _, item := range sale.SaleItems {
		var inventoryItem models.InventoryItem
		if err := tx.Where("product_id = ?", item.ProductID).First(&inventoryItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Produto não encontrado no estoque"})
			return
		}

		if inventoryItem.Quantity < item.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quantidade insuficiente em estoque"})
			return
		}

		// Atualiza quantidade
		inventoryItem.Quantity -= item.Quantity
		if err := tx.Save(&inventoryItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar estoque"})
			return
		}

		// Registra movimento de estoque
		movement := models.InventoryMovement{
			ProductID:     item.ProductID,
			MovementType:  "exit",
			Quantity:      -int(item.Quantity),
			PreviousStock: int(inventoryItem.Quantity) + int(item.Quantity),
			NewStock:      int(inventoryItem.Quantity),
			Reason:        "Venda #" + strconv.Itoa(int(sale.ID)),
			UserID:        sale.UserID,
		}
		if err := tx.Create(&movement).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar movimento de estoque"})
			return
		}
	}

	tx.Commit()

	// Recarrega a venda com os relacionamentos
	h.DB.Preload("Customer").Preload("User").Preload("SaleItems.Product").First(&sale, sale.ID)

	c.JSON(http.StatusCreated, gin.H{"sale": sale})
}

// UpdateSale atualiza uma venda existente
func (h *Handler) UpdateSale(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var sale models.Sale
	if err := h.DB.First(&sale, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Venda não encontrada"})
		return
	}

	var updateData models.Sale
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Recalcula valor final se necessário
	if updateData.TotalAmount > 0 || updateData.Discount >= 0 {
		totalAmount := updateData.TotalAmount
		if totalAmount == 0 {
			totalAmount = sale.TotalAmount
		}
		discount := updateData.Discount
		if updateData.Discount == 0 && sale.Discount != 0 {
			discount = sale.Discount
		}
		updateData.FinalAmount = totalAmount - discount
	}

	if err := h.DB.Model(&sale).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar venda"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sale": sale})
}

// GetSalesReport gera relatório de vendas
func (h *Handler) GetSalesReport(c *gin.Context) {
	var sales []models.Sale
	var totalSales float64
	var totalCount int64

	query := h.DB.Model(&models.Sale{})

	// Filtros opcionais
	if startDate := c.Query("start_date"); startDate != "" {
		query = query.Where("sale_date >= ?", startDate)
	}
	if endDate := c.Query("end_date"); endDate != "" {
		query = query.Where("sale_date <= ?", endDate)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Conta total e soma valores
	query.Count(&totalCount)
	query.Select("SUM(final_amount)").Row().Scan(&totalSales)

	// Busca vendas
	if err := query.Preload("Customer").Find(&sales).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar relatório"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sales":       sales,
		"total_sales": totalSales,
		"total_count": totalCount,
	})
}

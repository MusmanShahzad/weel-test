package handler
import (
	"net/http"
	"strconv"
	"weel-backend/internal/service"
	"github.com/gin-gonic/gin"
)
type OrderHandler struct {
	orderService service.OrderService
}
func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}
func (h *OrderHandler) RegisterRoutes(router *gin.RouterGroup) {
	orders := router.Group("/orders")
	{
		orders.GET("", h.GetOrders)
		orders.POST("", h.CreateOrder)
		orders.GET("/:id", h.GetOrder)
		orders.PUT("/:id", h.UpdateOrder)
	}
}
func (h *OrderHandler) GetAISuggestions(c *gin.Context) {
	var req service.GetAISuggestionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	suggestions, err := h.orderService.GetAISuggestions(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get AI suggestions"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"suggestions": suggestions,
		"count":       len(suggestions),
	})
}
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req service.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := h.orderService.CreateOrder(userID.(uint), &req)
	if err != nil {
		if err == service.ErrInvalidInput {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}
	c.JSON(http.StatusCreated, order)
}
func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var filters service.GetOrdersFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orders, err := h.orderService.GetOrders(userID.(uint), &filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"count":  len(orders),
	})
}
func (h *OrderHandler) GetOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}
	order, err := h.orderService.GetOrderByID(uint(id), userID.(uint))
	if err != nil {
		if err == service.ErrOrderNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == service.ErrUnauthorizedAccess {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get order"})
		return
	}
	c.JSON(http.StatusOK, order)
}
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}
	var req service.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := h.orderService.UpdateOrder(uint(id), userID.(uint), &req)
	if err != nil {
		if err == service.ErrOrderNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == service.ErrUnauthorizedAccess {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if err == service.ErrInvalidOrderStatus {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
		return
	}
	c.JSON(http.StatusOK, order)
}

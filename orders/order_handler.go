package orders

import (
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
	"shop/validation"
    "github.com/go-playground/validator/v10"
	"shop/users"
	"shop/products"
)

type OrderHandler struct {
    OrderRepo    *OrderRepository
    UserRepo     *users.UserRepository
    ProductRepo  *products.ProductRepository
}

func NewOrderHandler(or *OrderRepository, ur *users.UserRepository, pr *products.ProductRepository) *OrderHandler {
    return &OrderHandler{OrderRepo: or, UserRepo: ur, ProductRepo: pr}
}

func (oh *OrderHandler) CreateOrder(c *gin.Context) {
    var order Order
    if err := c.BindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
        return
    }

    if err := validation.ValidateStruct(&order); err != nil {
        errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), OrderBaseMessages)
        c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
        return
    }

	userIDStr := strconv.Itoa(int(order.UserID))
    if _, err := oh.UserRepo.GetUserByID(userIDStr); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    for _, productID := range order.ProductIDs {
        if _, err := oh.ProductRepo.GetProductByID(strconv.Itoa(int(productID))); err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product with ID " + strconv.Itoa(int(productID)) + " not found"})
            return
        }
    }

    if err := oh.OrderRepo.saveOrder(&order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving order"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully!"})
}

func (oh *OrderHandler) GetOrderByID(c *gin.Context) {
    id := c.Param("id")

    order, err := oh.OrderRepo.getOrderById(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (oh *OrderHandler) GetAllOrders(c *gin.Context) {
    orders, err := oh.OrderRepo.getAllOrders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving orders"})
        return
    }

	if len(orders) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "No orders found"})
        return
    }

    c.JSON(http.StatusOK, orders)
}

func (oh *OrderHandler) UpdateOrder(c *gin.Context) {
    id := c.Param("id")
    
	var updatedOrder Order
    if err := c.BindJSON(&updatedOrder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
        return
    }

    if err := validation.ValidateStruct(&updatedOrder); err != nil {
        errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), OrderBaseMessages)
        c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
        return
    }

	existingOrder, err := oh.OrderRepo.getOrderById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	updatedOrder.OrderDate = existingOrder.OrderDate

    if err := oh.OrderRepo.updateOrder(id, &updatedOrder); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully!"})
}

func (oh *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	_, err := oh.OrderRepo.getOrderById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

    if err := oh.OrderRepo.deleteOrder(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully!"})
}

func (oh *OrderHandler) SearchOrdersByUserID(c *gin.Context) {
	userID := c.Query("user")
    orders, err := oh.OrderRepo.searchOrdersByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching orders by user ID"})
        return
    }

    if len(orders) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "No orders found"})
        return
    }

    c.JSON(http.StatusOK, orders)
}

func (oh *OrderHandler) SearchOrdersByStatus(c *gin.Context) {
    status := c.Query("status")
    orders, err := oh.OrderRepo.searchOrdersByStatus(status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching orders by status"})
        return
    }

    if len(orders) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "No orders found for the given status"})
        return
    }

    c.JSON(http.StatusOK, orders)
}

package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"shop/products"
	"shop/users"
	"shop/orders"
	"shop/payments"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db := connectToDatabase(dbURL)
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error getting raw database object: %v\n", err)
		}
		sqlDB.Close()
	}()

	autoMigrate(db)

	productRepo := products.NewProductRepository(db)
	productHandler := products.NewProductHandler(productRepo)

	userRepo := users.NewUserRepository(db)
	userHandler := users.NewUserHandler(userRepo)

	orderRepo := orders.NewOrderRepository(db)
    orderHandler := orders.NewOrderHandler(orderRepo, userRepo, productRepo)

	paymentRepo := &payments.PaymentRepository{DB: db}
    paymentHandler := &payments.PaymentHandler{Repository: paymentRepo}

	router := gin.Default()

	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", productHandler.GetAllProducts)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.GetProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
		productRoutes.GET("/search", func(c *gin.Context) {
			if name := c.Query("name"); name != "" {
				productHandler.SearchProductsByName(c)
			} else if category := c.Query("category"); category != "" {
				productHandler.SearchProductsByCategory(c)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'name' or 'category' is required"})
			}
		})
	}

	userRoutes := router.Group("/users")
    {
        userRoutes.GET("/", userHandler.GetAllUsers)
        userRoutes.POST("/", userHandler.CreateUser)
        userRoutes.GET("/:id", userHandler.GetUserByID)
        userRoutes.PUT("/:id", userHandler.UpdateUser)
        userRoutes.DELETE("/:id", userHandler.DeleteUser)
        userRoutes.GET("/search", func(c *gin.Context) {
            if name := c.Query("name"); name != "" {
                userHandler.SearchUsersByName(c)
            } else if email := c.Query("email"); email != "" {
                userHandler.SearchUsersByEmail(c)
            } else {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'name' or 'email' is required"})
            }
        })
    }

	orderRoutes := router.Group("/orders")
    {
        orderRoutes.GET("/", orderHandler.GetAllOrders)
        orderRoutes.POST("/", orderHandler.CreateOrder)
        orderRoutes.GET("/:id", orderHandler.GetOrderByID)
        orderRoutes.PUT("/:id", orderHandler.UpdateOrder)
        orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
        orderRoutes.GET("/search", func(c *gin.Context) {
            if userID := c.Query("user"); userID != "" {
                orderHandler.SearchOrdersByUserID(c)
            } else if status := c.Query("status"); status != "" {
                orderHandler.SearchOrdersByStatus(c)
            } else {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'user' or 'status' is required"})
            }
        })
    }

	payments := router.Group("/payments")
	{
		payments.GET("/", paymentHandler.GetAllPayments)
		payments.POST("/", paymentHandler.CreatePayment)
		payments.GET("/:id", paymentHandler.GetPaymentByID)
		payments.PUT("/:id", paymentHandler.UpdatePayment)
		payments.DELETE("/:id", paymentHandler.DeletePayment)
		payments.GET("/search", func(c *gin.Context) {
			if userID := c.Query("user"); userID != "" {
				paymentHandler.SearchPaymentsByUserID(c)
			} else if orderID := c.Query("order"); orderID != "" {
				paymentHandler.SearchPaymentsByOrderID(c)
			} else if status := c.Query("status"); status != "" {
				paymentHandler.SearchPaymentsByStatus(c)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'user', 'order' or 'status' is required"})
			}
		})
	}

	port := ":8000"
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

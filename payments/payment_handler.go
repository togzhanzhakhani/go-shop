package payments

import (
	"net/http"
	"strconv"
	"log"
	"github.com/gin-gonic/gin"
)

const (
    DefaultPaymentData = `{
 "hpan":"4405639704015096","expDate":"0125","cvc":"815","terminalId":"67e34d63-102f-4bd1-898e-370781d0074d"
}`
)

type PaymentHandler struct {
    Repository *PaymentRepository
}

func (handler *PaymentHandler) GetAllPayments(c *gin.Context) {
    payments, err := handler.Repository.GetAllPayments()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payments)
}

func (handler *PaymentHandler) CreatePayment(c *gin.Context) {
    var payment Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := getPaymentToken()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	encryptedData, err := encryptData(DefaultPaymentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return
	}

    paymentResponse, err := makePayment(token, encryptedData)
    if err != nil {
		log.Printf("Failed to make payment: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make payment"})
        return
    }

    payment.PaymentStatus = paymentResponse.Status

    if err := handler.Repository.CreatePayment(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, payment)
}

func (handler *PaymentHandler) GetPaymentByID(c *gin.Context) {
    id := c.Param("id")
    payment, err := handler.Repository.GetPaymentByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (handler *PaymentHandler) UpdatePayment(c *gin.Context) {
    id := c.Param("id")
    var payment Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	payment.ID = uint(idUint)
	
    if err := handler.Repository.UpdatePayment(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (handler *PaymentHandler) DeletePayment(c *gin.Context) {
    id := c.Param("id")
    if err := handler.Repository.DeletePayment(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

func (handler *PaymentHandler) SearchPaymentsByUserID(c *gin.Context) {
    userID := c.Query("user")
    payments, err := handler.Repository.SearchPaymentsByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payments)
}

func (handler *PaymentHandler) SearchPaymentsByOrderID(c *gin.Context) {
    orderID := c.Query("order")
    payments, err := handler.Repository.SearchPaymentsByOrderID(orderID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payments)
}

func (handler *PaymentHandler) SearchPaymentsByStatus(c *gin.Context) {
    status := c.Query("status")
    payments, err := handler.Repository.SearchPaymentsByStatus(status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payments)
}
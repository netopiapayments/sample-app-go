package handlers

import (
	"net/http"

	"sample/services"

	"github.com/gin-gonic/gin"
	"github.com/netopiapayments/go-sdk/requests"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
}

func NewPaymentHandler(ps *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: ps}
}

// POST /start-payment
func (h *PaymentHandler) StartPayment(c *gin.Context) {
	var startReq requests.StartPaymentRequest
	if err := c.ShouldBindJSON(&startReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	resp, err := h.paymentService.StartPayment(&startReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start payment", "details": err.Error()})
		return
	}

	if resp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No response from payment service"})
		return
	}

	if resp.Payment == nil {
		c.JSON(http.StatusOK, gin.H{"message": resp.Message})
		return
	}

	if resp.Error != nil && resp.Error.Code != "00" {
		c.JSON(http.StatusOK, gin.H{
			"error":          "Payment initiation failed",
			"error_message":  resp.Error.Message,
			"error_code":     resp.Error.Code,
			"payment_url":    resp.Payment.PaymentURL,
			"payment_token":  resp.Payment.Token,
			"payment_status": resp.Payment.Status,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Payment initiated successfully",
		"payment_url":    resp.Payment.PaymentURL,
		"payment_token":  resp.Payment.Token,
		"payment_ntpId":  resp.Payment.NtpID,
		"payment_status": resp.Payment.Status,
	})
}

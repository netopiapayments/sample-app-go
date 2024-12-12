package handlers

import (
	"net/http"

	"sample/services"

	"github.com/gin-gonic/gin"
)

type StatusPayload struct {
	NtpID   string `json:"ntpID"`
	OrderID string `json:"orderID"`
}

type StatusHandler struct {
	paymentService *services.PaymentService
}

func NewStatusHandler(ps *services.PaymentService) *StatusHandler {
	return &StatusHandler{paymentService: ps}
}

// POST /status
func (h *StatusHandler) GetStatus(c *gin.Context) {
	var payload StatusPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	if payload.NtpID == "" || payload.OrderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ntpID and orderID are required"})
		return
	}

	statusResp, err := h.paymentService.GetStatus(payload.NtpID, payload.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get status", "details": err.Error()})
		return
	}

	if statusResp == nil || statusResp.Payment == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No payment information found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code":    statusResp.Payment.Status,
		"message":        statusResp.Message,
		"payment_method": statusResp.Payment.Method,
		"amount":         statusResp.Payment.Amount,
		"currency":       statusResp.Payment.Currency,
	})
}

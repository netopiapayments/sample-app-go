package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /ipn
func (h *PaymentHandler) IPNHandler(c *gin.Context) {
	verificationResult, err := h.paymentService.VerifyIPN(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IPN Verification failed", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "IPN processed successfully",
		"status":  verificationResult.Status,
		"info":    verificationResult.Message,
	})
}

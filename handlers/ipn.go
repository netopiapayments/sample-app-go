package handlers

import (
	"net/http"
	"sample/services"

	"github.com/gin-gonic/gin"
)

type IPNHandler struct {
	paymentService *services.NetopiaService
}

func NewIPNHandler(ps *services.NetopiaService) *IPNHandler {
	return &IPNHandler{paymentService: ps}
}

// POST /ipn
func (h *IPNHandler) IPNHandler(c *gin.Context) {
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

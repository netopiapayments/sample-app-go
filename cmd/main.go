package main

import (
	"log"

	"sample/config"
	"sample/handlers"
	"sample/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	netopiaClient, err := config.NewNetopiaClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Netopia client: %v", err)
	}

	paymentService := services.NewPaymentService(netopiaClient, cfg)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	statusHandler := handlers.NewStatusHandler(paymentService)

	r := gin.Default()

	r.POST("/start-payment", paymentHandler.StartPayment)
	r.POST("/ipn", paymentHandler.IPNHandler)
	r.POST("/status", statusHandler.GetStatus)

	r.Run(":8080")
}

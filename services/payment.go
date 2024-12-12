package services

import (
	"fmt"
	"net/http"

	"sample/config"

	netopia "github.com/netopiapayments/go-sdk"
	"github.com/netopiapayments/go-sdk/requests"
	"github.com/netopiapayments/go-sdk/responses"
)

type PaymentService struct {
	client *netopia.PaymentClient
	cfg    *config.Config
}

func NewPaymentService(c *netopia.PaymentClient, cfg *config.Config) *PaymentService {
	return &PaymentService{client: c, cfg: cfg}
}

func (ps *PaymentService) StartPayment(req *requests.StartPaymentRequest) (*responses.StartPaymentResponse, error) {
	req.Order.PosSignature = ps.cfg.PosSignature
	req.Config.NotifyURL = ps.cfg.NotifyURL
	req.Config.RedirectURL = ps.cfg.RedirectURL

	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return ps.client.StartPayment(req)
}

func (ps *PaymentService) VerifyIPN(r *http.Request) (*netopia.IPNVerificationResult, error) {
	return ps.client.VerifyIPN(r)
}

func (ps *PaymentService) GetStatus(ntpID, orderID string) (*responses.StatusResponse, error) {
	return ps.client.GetStatus(ntpID, orderID)
}

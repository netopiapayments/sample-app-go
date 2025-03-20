package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	netopia "github.com/netopiapayments/go-sdk"
)

type Config struct {
	APIKey          string
	PosSignature    string
	PublicKey       []byte
	IsLive          bool
	NotifyURL       string
	RedirectURL     string
	PosSignatureSet []string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	isLiveStr := os.Getenv("IS_LIVE")
	isLive, err := strconv.ParseBool(isLiveStr)
	if err != nil {
		isLive = false
	}

	cfg := &Config{
		APIKey:       os.Getenv("API_KEY"),
		PosSignature: os.Getenv("POS_SIGNATURE"),
		PublicKey:    []byte(os.Getenv("PUBLIC_KEY")),
		IsLive:       isLive,
		NotifyURL:    os.Getenv("NOTIFY_URL"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
	}

	cfg.PosSignatureSet = []string{cfg.PosSignature}
	return cfg, nil
}

func NewNetopiaClient(cfg *Config) (*netopia.PaymentClient, error) {
	netCfg := netopia.Config{
		ApiKey:          cfg.APIKey,
		PosSignature:    cfg.PosSignature,
		IsLive:          cfg.IsLive,
		NotifyURL:       cfg.NotifyURL,
		RedirectURL:     cfg.RedirectURL,
		PublicKey:       cfg.PublicKey,
		PosSignatureSet: cfg.PosSignatureSet,
	}

	logger := &netopia.DefaultLogger{}
	return netopia.NewPaymentClient(netCfg, logger)
}

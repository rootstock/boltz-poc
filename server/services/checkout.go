package services

import (
	"github.com/btcsuite/btcutil"
	"github.com/patogallaiov/boltz-poc/config"
	"github.com/patogallaiov/boltz-poc/connectors"
	"github.com/patogallaiov/boltz-poc/storage"
	log "github.com/sirupsen/logrus"
)

type Invoice struct {
	Lninvoice string `json:"lninvoice"`
}

type CheckoutService struct {
	boltz connectors.BoltzConnector
	db    storage.DBConnector
}

func NewCheckoutService(appCfg config.Config, boltz connectors.BoltzConnector, db storage.DBConnector) *CheckoutService {
	return &CheckoutService{
		boltz,
		db,
	}
}

func (service *CheckoutService) CreateInvoice(request storage.PaymentRequest) (*Invoice, error) {
	swaptype, err := service.db.GetConfig("swaptype")
	if err != nil {
		return &Invoice{}, err
	}
	log.Debug("Swap type config:", swaptype.Value)
	pair := "BTC/rBTC"
	switch swaptype.Value {
	case "liquidity":
		pair = "BTC/DOC"
	case "mint", "default":
		pair = "BTC/rBTC"
	}

	response, err := service.boltz.NewReverseSwap(pair, "sell", btcutil.Amount(request.ProductId*2000), "", nil)
	if err != nil {
		return &Invoice{}, err
	}
	service.db.SavePayment(&storage.Payment{
		PreimageHash: response.PreimageHash,
		Preimage:     response.Preimage,
		Invoice:      response.Invoice,
		Amount:       request.ProductId * 2000,
		ProductId:    request.ProductId,
		Status:       "New",
		Tx:           "-",
	})
	return &Invoice{
		Lninvoice: response.Invoice,
	}, nil
}

func (service *CheckoutService) GetInvoices() (any, error) {
	return service.db.GetPayments()
}

func (service *CheckoutService) GetInvoice(preimageHash string) (any, error) {
	return service.db.GetPayment(preimageHash)
}

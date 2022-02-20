package services

import (
	"github.com/Hitsugaya/rest-api-project/cmd/internal/config"
)

type ExchangeService interface {
	GetRate(coordX string, coordY string) (float64, error)
}
type ExternalExchangeService interface {
	GetDollarExchangeRate() (float64, error)
	GetEuroExchangeRate() (float64, error)
}

type AllServices struct {
	ExchangeService
	ExternalExchangeService
}

func NewAllServices(circumferenceConfig *config.Circumference) *AllServices {
	eEService := NewExternalExchangeService()
	return &AllServices{
		ExternalExchangeService: eEService,
		ExchangeService:         NewExchangeService(circumferenceConfig, eEService),
	}
}

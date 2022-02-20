package services

import (
	"github.com/Hitsugaya/rest-api-project/cmd/internal/config"
	"math"
	"strconv"
)

var _ ExchangeService = &exchangeService{}

type exchangeService struct {
	circumferenceConfig     *config.Circumference
	externalExchangeService *externalExchangeService
}

func (e *exchangeService) GetRate(coordXStr string, coordYStr string) (float64, error) {
	coordX, err := strconv.ParseFloat(coordXStr, 64)
	if err != nil {
		return 0, err
	}
	coordY, err := strconv.ParseFloat(coordYStr, 64)
	if err != nil {
		return 0, err
	}
	f := math.Pow(coordX-e.circumferenceConfig.CenterCoordinateX, 2) +
		math.Pow(coordY-e.circumferenceConfig.CenterCoordinateY, 2)

	R := math.Pow(e.circumferenceConfig.Diameter/2, 2)
	if f <= R {
		return e.externalExchangeService.GetDollarExchangeRate()
	} else {
		return e.externalExchangeService.GetEuroExchangeRate()
	}
}

func NewExchangeService(circumferenceConfig *config.Circumference, eEs *externalExchangeService) *exchangeService {
	return &exchangeService{
		circumferenceConfig:     circumferenceConfig,
		externalExchangeService: eEs,
	}
}

package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var _ ExternalExchangeService = &externalExchangeService{}

const url = "https://api.exchangerate.host/latest?base=%s&symbols=%s"

type externalExchangeService struct {
}

func (e *externalExchangeService) GetDollarExchangeRate() (float64, error) {
	client := http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf(url, "USD", "RUB"), nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		return 0, nil
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["rates"].(map[string]interface{})["RUB"].(float64), nil
}

func (e *externalExchangeService) GetEuroExchangeRate() (float64, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf(url, "EUR", "RUB"), nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		return 0, nil
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["rates"].(map[string]interface{})["RUB"].(float64), nil
}

func NewExternalExchangeService() *externalExchangeService {
	return &externalExchangeService{}
}

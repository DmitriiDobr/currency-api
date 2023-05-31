package exchangerate

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type CurrenciesExchangeRate struct {
	*Parser
}

func NewCurrenciesExchangeRate() *CurrenciesExchangeRate {
	return &CurrenciesExchangeRate{NewParser(http.Client{})}
}

type Parser struct {
	client *http.Client
	result map[string]interface{}
}

func NewParser(client http.Client) *Parser {
	return &Parser{client: &client}
}

func (p *Parser) createRequest(baseCurrency, quoteCurrency string) (*http.Request, error) {
	url := "https://api.exchangerate.host/latest?base=" + baseCurrency + "&symbols=" + quoteCurrency
	return http.NewRequest("GET", url, nil)

}

func (p *Parser) doRequest(r *http.Request) (*http.Response, error) {
	return p.client.Do(r)
}

func (p *Parser) decodeResponse(resp *http.Response) {
	json.NewDecoder(resp.Body).Decode(&p.result)
}

func (c *CurrenciesExchangeRate) GetCurrencyRate(baseCurrency, quoteCurrency string) (float64, error) {
	request, err := c.createRequest(baseCurrency, quoteCurrency)
	if err != nil {
		log.Fatal("Не получилось создать запрос")
		return 0, err
	}
	response, err := c.doRequest(request)
	if err != nil {
		log.Fatal("Не получилось сделать запрос")
		return 0, err
	}
	c.decodeResponse(response)
	rateLevel := c.result["rates"]
	if exchangeRate, ok := rateLevel.(map[string]interface{})[quoteCurrency].(float64); ok {
		return exchangeRate, nil
	}
	return 0, errors.New("Тип данных не соответствует float64")

}

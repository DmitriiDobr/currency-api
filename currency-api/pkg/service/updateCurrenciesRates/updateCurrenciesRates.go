package updateCurrenciesRates

import (
	"encoding/json"
	"net/http"
)

type CurrencyExchangeRateParser interface {
	GetCurrencyRate(baseCurrency, quoteCurrency string) (float64, error)
}

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
	request, err := http.NewRequest("GET", url, nil)
	return request, err

}

func (p *Parser) doRequest(r *http.Request) (*http.Response, error) {
	resp, err := p.client.Do(r)
	return resp, err
}

func (p *Parser) decodeResponse(resp *http.Response) {
	json.NewDecoder(resp.Body).Decode(&p.result)
}

func (c *CurrenciesExchangeRate) GetCurrencyRate(baseCurrency, quoteCurrency string) (float64, error) {
	request, err := c.createRequest(baseCurrency, quoteCurrency)
	if err != nil {
		return 0, err
	}
	response, err := c.doRequest(request)
	if err != nil {
		return 0, err
	}
	c.decodeResponse(response)
	rateLevel := c.result["rates"]
	exchangeRate := rateLevel.(map[string]interface{})[quoteCurrency].(float64)
	return exchangeRate, err
}

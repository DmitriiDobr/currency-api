package repository

import "time"

type CurrencyPairAnswer struct {
}

type CurrencyPair struct {
	Id           int       `json:"id" db:"id"`
	CurrencyFrom string    `json:"CurrencyFrom" db:"currencyfrom"`
	CurrencyTo   string    `json:"CurrencyTo" db:"currencyto"`
	ExchangeRate float64   `json:"ExchangeRate" db:"exchangerate"`
	UpdatedAt    time.Time `json:"UpdatedAt" db:"updatedat"`
}

type AddCurrencyPair struct {
	CurrencyFrom *string `json:"CurrencyFrom"`
	CurrencyTo   *string `json:"CurrencyTo"`
}

type Currencies struct {
	ID           int    `json:"ID"`
	CurrencyFrom string `json:"CurrencyFrom"`
	CurrencyTo   string `json:"CurrencyTo"`
}

type CurrencyPairs struct {
	C []Currencies
}

type ConversionOfMoney struct {
	CurrencyFrom string `json:"CurrencyFrom"`
	CurrencyTo   string `json:"CurrencyTo"`
	Amount       string `json:"Amount"`
}

type ConvertedAmount struct {
	ConversionRate string `json:"ConversionRate"`
}

type CurrencyInsert struct {
	CurrencyFrom string    `json:"CurrencyFrom" db:"currencyfrom"`
	CurrencyTo   string    `json:"CurrencyTo" db:"currencyto"`
	ExchangeRate float64   `json:"ExchangeRate" db:"exchangerate"`
	UpdatedAt    time.Time `json:"UpdatedAt" db:"updatedat"`
}

package repository

import "time"

type CurrencyPairAnswer struct {
}

type CurrencyPair struct {
	Id           int       `json:"id" db:"id"`
	CurrencyFrom string    `json:"currency_from" db:"currency_from"`
	CurrencyTo   string    `json:"currency_to" db:"currency_to"`
	ExchangeRate float64   `json:"exchange_rate" db:"exchange_rate"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type AddCurrencyPair struct {
	CurrencyFrom *string `json:"currency_from"`
	CurrencyTo   *string `json:"currency_to"`
}

type Currencies struct {
	ID           int    `json:"id" db:"id"`
	CurrencyFrom string `json:"currency_from" db:"currency_from"`
	CurrencyTo   string `json:"currency_to" db:"currency_to"`
}

type CurrencyPairs struct {
	C []Currencies
}

type ConversionOfMoney struct {
	CurrencyFrom string `json:"currency_from"`
	CurrencyTo   string `json:"currency_to"`
	Amount       string `json:"amount"`
}

type ConvertedAmount struct {
	ConversionRate string `json:"conversion_rate"`
}

type CurrencyInsert struct {
	CurrencyFrom string    `json:"currency_from" db:"currency_from"`
	CurrencyTo   string    `json:"currency_to" db:"currency_to"`
	ExchangeRate float64   `json:"exchange_rate" db:"exchange_rate"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

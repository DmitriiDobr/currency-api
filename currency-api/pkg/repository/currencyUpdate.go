package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type updateCurrencyDb interface {
	UpdateExchangeRates(baseCurrencies, quoteCurrencies string, exchangeRates float64) error
	GetAllPairsData() ([]CurrencyPair, error)
}

type UpdateCurrencyDb struct {
	db *sqlx.DB
	updateCurrencyDb
}

func NewUpdateCurrencyDb(db *sqlx.DB) *UpdateCurrencyDb {
	return &UpdateCurrencyDb{db: db}
}

func (u *UpdateCurrencyDb) UpdateExchangeRates(baseCurrencies, quoteCurrencies string, exchangeRates float64) error {

	query := fmt.Sprintf(`
UPDATE 
    currency_exchange 
SET exchangerate = $1,updatedat=$2 
WHERE currencyfrom=$3 and currencyto=$4 `)
	_, err := u.db.Exec(query, exchangeRates, time.Now(), baseCurrencies, quoteCurrencies)
	return err
}

func (u *UpdateCurrencyDb) GetAllPairsData() ([]CurrencyPair, error) {
	var pairsList []CurrencyPair
	query := fmt.Sprintf(`
SELECT 
    cur.id, cur.CurrencyFrom, cur.CurrencyTo,cur.ExchangeRate,cur.UpdatedAt 
FROM currency_exchange cur`)
	if err := u.db.Select(&pairsList, query); err != nil {
		return nil, err
	}
	return pairsList, nil

}

package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type updateCurrencyDb interface {
	UpdateExchangeRates(ctx context.Context, baseCurrencies, quoteCurrencies string, exchangeRates float64) error
	GetAllPairsData(ctx context.Context) ([]CurrencyPair, error)
}

type UpdateCurrencyDb struct {
	db *sqlx.DB
}

func NewUpdateCurrencyDb(db *sqlx.DB) *UpdateCurrencyDb {
	return &UpdateCurrencyDb{db: db}
}

func (u *UpdateCurrencyDb) UpdateExchangeRates(ctx context.Context, baseCurrencies, quoteCurrencies string, exchangeRates float64) error {
	query := fmt.Sprintf(`
UPDATE 
    currency_exchange 
SET exchange_rate = $1,updated_at=$2 
WHERE currency_from=$3 and currency_to=$4 `)
	_, err := u.db.ExecContext(ctx, query, exchangeRates, time.Now(), baseCurrencies, quoteCurrencies)
	return err
}

func (u *UpdateCurrencyDb) GetAllPairsData(ctx context.Context) ([]CurrencyPair, error) {
	var pairsList []CurrencyPair
	query := fmt.Sprintf(`
SELECT 
    cur.id, cur.currency_from, cur.currency_to,cur.exchange_rate,cur.updated_at 
FROM currency_exchange cur`)
	if err := u.db.SelectContext(ctx, &pairsList, query); err != nil {
		return nil, err
	}
	return pairsList, nil

}

package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type currenciesInfo interface {
	GetAllPairs(ctx context.Context) ([]Currencies, error)
	ConvertMoney(ctx context.Context, baseCurrency, quoteCurrency string, value float64) (float64, error)
	AddNewPair(ctx context.Context, baseCurrency, quoteCurrency string, exchangeRate float64) (int, error)
}

type CurrencyDb struct {
	db *sqlx.DB
	currenciesInfo
}

func NewCurrencyDb(db *sqlx.DB) *CurrencyDb {
	return &CurrencyDb{db: db}
}

func (c *CurrencyDb) GetAllPairs(ctx context.Context) (pairs []Currencies, err error) {
	//передача контекста
	query := fmt.Sprintf(`
SELECT 
	cur.id, cur.CurrencyFrom, cur.CurrencyTo 
from currency_exchange cur`)
	return pairs, c.db.SelectContext(
		ctx,
		&pairs,
		query,
	)
}

func (c *CurrencyDb) ConvertMoney(ctx context.Context, baseCurrency, quoteCurrency string, amount float64) (exchangeRate float64, err error) {
	query := fmt.Sprintf(`
SELECT 
    cur.ExchangeRate 
FROM currency_exchange cur
WHERE currencyFrom=$1 and currencyTo=$2
       `)
	return exchangeRate * amount, c.db.GetContext(ctx,
		&exchangeRate,
		query,
		baseCurrency,
		quoteCurrency,
	)

}

func (c *CurrencyDb) AddNewPair(ctx context.Context, baseCurrency, quoteCurrency string, exchangeRate float64) (idCurrency int, err error) {

	insert := CurrencyInsert{
		baseCurrency,
		quoteCurrency,
		exchangeRate,
		time.Now(),
	}

	query := `
INSERT into currency_exchange 
    (CurrencyFrom,CurrencyTo,ExchangeRate,UpdatedAt)																			
values (:currencyfrom,:currencyto,:exchangerate,:updatedat) RETURNING id`

	query, args, err := sqlx.Named(query, insert)
	query = c.db.Rebind(query)
	return idCurrency, c.db.GetContext(ctx, &idCurrency, query, args...)

}

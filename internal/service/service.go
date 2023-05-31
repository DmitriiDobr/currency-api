package service

import (
	"context"
	"currencyapi/internal/repository"
	"currencyapi/pkg/exchangerate"
	"fmt"
	"log"
)

type currencyManipulation interface {
	AddNewPair(ctx context.Context, baseCurrency, quoteCurrency string) (int, error)
	GetAllPairs(ctx context.Context) ([]repository.Currencies, error)
	ConvertMoney(ctx context.Context, baseCurrency, quoteCurrency string, amount float64) (float64, error)
}

type Service struct {
	*repository.CurrencyDb
	update *exchangerate.CurrenciesExchangeRate
	currencyManipulation
}

func NewService(repo *repository.CurrencyDb) *Service {
	return &Service{CurrencyDb: repo, update: exchangerate.NewCurrenciesExchangeRate()}
}

func (s *Service) AddNewPair(ctx context.Context, baseCurrency, quoteCurrency string) (int, error) {
	exchangeRate, err := s.update.GetCurrencyRate(baseCurrency, quoteCurrency)
	fmt.Println(exchangeRate)
	if err != nil {
		log.Fatal("Не удалось получить обменный курс!")
		return 0, nil
	}
	return s.CurrencyDb.AddNewPair(ctx, baseCurrency, quoteCurrency, exchangeRate)

}

func (s *Service) GetAllPairs(ctx context.Context) ([]repository.Currencies, error) {
	return s.CurrencyDb.GetAllPairs(ctx)
}

func (s *Service) ConvertMoney(ctx context.Context, baseCurrency, quoteCurrency string, amount float64) (float64, error) {
	return s.CurrencyDb.ConvertMoney(ctx, baseCurrency, quoteCurrency, amount)

}

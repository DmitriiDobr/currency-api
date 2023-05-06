package service

import (
	"ConverterService/pkg/repository"
	"ConverterService/pkg/service/updateCurrenciesRates"
	"context"
	"fmt"
)

type currencyManipulation interface {
	AddNewPair(ctx context.Context, baseCurrency, quoteCurrency string) (int, error)
	GetAllPairs(ctx context.Context) ([]repository.Currencies, error)
	ConvertMoney(ctx context.Context, baseCurrency, quoteCurrency string, amount float64) (float64, error)
}

type Service struct {
	*repository.CurrencyDb
	update *updateCurrenciesRates.CurrenciesExchangeRate
	currencyManipulation
}

func NewService(repo *repository.CurrencyDb) *Service {
	return &Service{CurrencyDb: repo, update: updateCurrenciesRates.NewCurrenciesExchangeRate()}
}

func (s *Service) AddNewPair(ctx context.Context, baseCurrency, quoteCurrency string) (int, error) {
	exchangeRate, err := s.update.GetCurrencyRate(baseCurrency, quoteCurrency)
	fmt.Println(exchangeRate)
	if err != nil {
		return 0, nil
	}
	id, err := s.CurrencyDb.AddNewPair(ctx, baseCurrency, quoteCurrency, exchangeRate)
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (s *Service) GetAllPairs(ctx context.Context) ([]repository.Currencies, error) {
	pairs, err := s.CurrencyDb.GetAllPairs(ctx)
	if err != nil {
		return nil, err
	}
	return pairs, nil
}

func (s *Service) ConvertMoney(ctx context.Context, baseCurrency, quoteCurrency string, amount float64) (float64, error) {
	money, err := s.CurrencyDb.ConvertMoney(ctx, baseCurrency, quoteCurrency, amount)
	if err != nil {
		return 0, err
	}
	fmt.Println("money", money)
	return money, nil

}

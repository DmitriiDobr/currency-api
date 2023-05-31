package worker

import (
	"context"
	"currencyapi/internal/repository"
	"currencyapi/pkg/exchangerate"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func ScheduleUpdates(db *sqlx.DB, ctx context.Context, timeS int) {
	go func() {
		for range time.Tick(time.Duration(timeS) * time.Second) {
			curDb := repository.NewUpdateCurrencyDb(db)

			pairs, err := curDb.GetAllPairsData(ctx)
			if err != nil {
				return
			}
			exchangeRate := exchangerate.NewCurrenciesExchangeRate()
			for _, pair := range pairs {
				pair := pair

				func() {
					exchangeRates, err := exchangeRate.GetCurrencyRate(pair.CurrencyFrom, pair.CurrencyTo)
					if err != nil {
						log.Fatal("")
						return
					}
					err = curDb.UpdateExchangeRates(ctx, pair.CurrencyFrom, pair.CurrencyTo, exchangeRates)
					if err != nil {
						return
					}
					fmt.Printf("Все обновилось - базовая валюта: %s, валюта котировки: %s, курс обмена: %v\n",
						pair.CurrencyFrom, pair.CurrencyTo, exchangeRates)

				}()

			}

		}
	}()
}

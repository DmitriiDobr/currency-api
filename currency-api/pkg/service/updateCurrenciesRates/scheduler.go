package updateCurrenciesRates

import (
	"ConverterService/pkg/repository"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

func ScheduleUpdates(db *sqlx.DB) {
	go func() {
		for range time.Tick(time.Second * 5) {
			curDb := repository.NewUpdateCurrencyDb(db)

			pairs, err := curDb.GetAllPairsData()
			if err != nil {
				return
			}
			exchangeRate := NewCurrenciesExchangeRate()
			var wg sync.WaitGroup
			var mutex sync.Mutex
			for _, pair := range pairs {
				pair := pair
				wg.Add(1)

				go func() {
					mutex.Lock()
					defer wg.Done()
					exchangeRates, err := exchangeRate.GetCurrencyRate(pair.CurrencyFrom, pair.CurrencyTo)
					if err != nil {
						return
					}
					err = curDb.UpdateExchangeRates(pair.CurrencyFrom, pair.CurrencyTo, exchangeRates)
					if err != nil {
						return
					}
					fmt.Printf("Все обновилось - базовая валюта: %s, валюта котировки: %s, курс обмена: %v\n",
						pair.CurrencyFrom, pair.CurrencyTo, exchangeRates)
					mutex.Unlock()

				}()

			}
			wg.Wait()
		}
	}()
}

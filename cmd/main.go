package main

import (
	"context"
	"currencyapi/internal/handlers"
	"currencyapi/internal/repository"
	"currencyapi/internal/service"
	"currencyapi/pkg/worker"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	ctx := context.Background()
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	app := fiber.New()
	fmt.Println("Staring server...")

	conf := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DbName:   viper.GetString("db.dbName"),
		SslMode:  viper.GetString("db.sslmode"),
	}
	conn, err := conf.InitDb()
	db := repository.NewCurrencyDb(conn)
	serviceLogic := service.NewService(db)
	routing := handlers.NewHandler(serviceLogic)
	routing.RegisterHandlers(app)

	if timeSeconds, err := strconv.Atoi(viper.GetString("time_seconds")); err == nil {
		worker.ScheduleUpdates(conn, ctx, timeSeconds)
	} else {
		log.Fatal("Неправильный конфиг!")
	}

	err = app.Listen(viper.GetString("address"))
	if err != nil {
		panic(err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

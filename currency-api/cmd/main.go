package main

import (
	"ConverterService/pkg/handlers"
	"ConverterService/pkg/repository"
	"ConverterService/pkg/service"
	"ConverterService/pkg/service/updateCurrenciesRates"
	"ConverterService/server"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	app := server.NewFiberServer()
	fmt.Println("Staring server...")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(viper.GetString("port"))
	})

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
	updateCurrenciesRates.ScheduleUpdates(conn)

	err = app.Listen(viper.GetString("address"))
	if err != nil {
		return
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

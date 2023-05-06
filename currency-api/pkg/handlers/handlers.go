package handlers

import (
	"ConverterService/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterHandlers(app *fiber.App) {

	api := app.Group("/api/currency")
	api.Get("/", h.getAllPairs)
	api.Put("/", h.convertMoney)
	api.Post("/", h.addNewPair)

}

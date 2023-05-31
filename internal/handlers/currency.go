package handlers

import (
	"currencyapi/internal/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) addNewPair(ctx *fiber.Ctx) error {
	body := repository.AddCurrencyPair{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if _, err := h.service.AddNewPair(ctx.Context(), *body.CurrencyTo, *body.CurrencyFrom); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(&body)

}

func (h *Handler) getAllPairs(ctx *fiber.Ctx) error {
	pairs, err := h.service.GetAllPairs(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&pairs)
}

func (h *Handler) convertMoney(ctx *fiber.Ctx) error {
	body := repository.ConversionOfMoney{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	amount, err := strconv.ParseFloat(body.Amount, 64)

	money, err := h.service.ConvertMoney(ctx.Context(), body.CurrencyFrom, body.CurrencyTo, amount)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	res := repository.ConvertedAmount{ConversionRate: strconv.FormatFloat(money, 'f', 6, 64)}
	return ctx.Status(fiber.StatusOK).JSON(&res)
}

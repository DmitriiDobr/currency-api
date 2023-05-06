package server

import "github.com/gofiber/fiber/v2"

func NewFiberServer() *fiber.App {
	return fiber.New()
}

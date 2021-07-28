package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hugolhafner/wendigo/internal/env"
	collect "github.com/hugolhafner/wendigo/internal/server/routes"
	"github.com/hugolhafner/wendigo/internal/services/pulsar"
	"sync"
)

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		pulsar.Init()
	})
}

func main() {
	Init()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	collector := app.Group("/collect")
	collector.Post("/mm", collect.MouseMovements)
	collector.Post("/fp", collect.Fingerprint)

	app.Listen(fmt.Sprintf("0.0.0.0:%d", env.PORT))
}

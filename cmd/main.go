package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hugolhafner/wendigo/internal/env"
	collect "github.com/hugolhafner/wendigo/internal/server/routes"
	"github.com/hugolhafner/wendigo/internal/services/kafka"
	"github.com/sirupsen/logrus"
)

var initOnce sync.Once
var shutdownOnce sync.Once

func Init() {
	initOnce.Do(func() {
		kafka.Init()
	})
}

func Shutdown() {
	shutdownOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		wg := &sync.WaitGroup{}

		go kafka.Shutdown(ctx, wg)

		wg.Wait()
	})
}

func main() {
	Init()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Use(logger.New(logger.Config{Format: "${ip} - ${latency} [${time}] \"${path}\" ${status} ${bytesSent} \"${referer}\" \"${ua}\"\n"}))
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	collector := app.Group("/collector")
	collector.Post("/mm", collect.MouseMovements)
	collector.Post("/fp", collect.Fingerprint)

	go func() {
		// Listen on application shutdown signals.
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

		// Start waiting for the signal, blocking until received
		<-listener

		logrus.Info("Starting graceful shutdown")

		// If this isn't hosted on kubernetes add extra cancel system as grace period
		_, cancel := context.WithTimeout(context.Background(), time.Second*35) // 35, default k8s is 30 so will sigkill before this
		defer cancel()

		// Shutdown HTTP server.
		if err := app.Shutdown(); err != nil {
			logrus.Warnf("App error on shutdown: %v", err)
			return
		}

		Shutdown()

		logrus.Infof("Successfully shutdown cleanly")
	}()

	app.Listen(fmt.Sprintf("0.0.0.0:%d", env.PORT))
}

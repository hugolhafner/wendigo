package collect

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hugolhafner/wendigo/internal/services/pulsar"
	"github.com/sirupsen/logrus"
)

func Fingerprint(ctx *fiber.Ctx) error {
	if err := pulsar.Publish.Fingerprint(ctx.Context()); err != nil {
		logrus.Warnf("Failed to publish fingerprint %v", "FINGERPRINT_ID")
		return err
	}

	return nil
}

func MouseMovements(ctx  *fiber.Ctx) error {
	if err := pulsar.Publish.MouseMovements(ctx.Context()); err != nil {
		logrus.Warnf("Failed to publish mouse movements %v", "OTHER_ID")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
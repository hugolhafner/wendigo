package collect

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/hugolhafner/wendigo/internal/data"
	"github.com/hugolhafner/wendigo/internal/services/pulsar"
	"github.com/sirupsen/logrus"
	"github.com/twmb/murmur3"
)

func Fingerprint(ctx *fiber.Ctx) error {
	fingerprint := data.Fingerprint{}

	if err := ctx.BodyParser(&fingerprint); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if isValid := data.ValidateStruct(fingerprint); !isValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]bool{"success": false})
	}

	data, err := json.Marshal(fingerprint)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]bool{"success": false})
	}

	uniqueHash := murmur3.Sum64(data)
	if err := pulsar.Publish.Fingerprint(ctx.Context(), uniqueHash, &data); err != nil {
		logrus.Warnf("Failed to publish fingerprint %v", "FINGERPRINT_ID")
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]bool{"success": true})
}

func MouseMovements(ctx *fiber.Ctx) error {
	if err := pulsar.Publish.MouseMovements(ctx.Context()); err != nil {
		logrus.Warnf("Failed to publish mouse movements %v", "OTHER_ID")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

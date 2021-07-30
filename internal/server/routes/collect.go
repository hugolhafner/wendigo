package collect

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/hugolhafner/wendigo/internal/data"
	"github.com/hugolhafner/wendigo/internal/services/kafka"
	"github.com/sirupsen/logrus"
	"github.com/twmb/murmur3"
)

func Fingerprint(ctx *fiber.Ctx) error {
	fingerprint := data.Fingerprint{}

	if err := ctx.BodyParser(&fingerprint); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}

	if isValid := data.ValidateStruct(fingerprint); !isValid {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	data, err := json.Marshal(fingerprint)
	if err != nil {
		return err
	}

	uniqueHash := murmur3.Sum64(data)
	if err := kafka.Publish.Fingerprint(ctx.Context(), uniqueHash, &data); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]bool{"success": true})
}

func MouseMovements(ctx *fiber.Ctx) error {
	movements := data.MouseMovements{}

	if err := ctx.BodyParser(&movements); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if isValid := data.ValidateStruct(movements); !isValid {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	data, err := json.Marshal(movements)
	if err != nil {
		return err
	}

	if err := kafka.Publish.MouseMovements(ctx.Context(), &data); err != nil {
		logrus.Warnf("Failed to publish mouse movements %v", "OTHER_ID")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return nil
}

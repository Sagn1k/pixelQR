package handlers

import (
    "github.com/Sagn1k/pixelQR/internal/qr"
    "github.com/Sagn1k/pixelQR/pkg/models"
    "go.uber.org/zap"
    "github.com/gofiber/fiber/v2"
)

func GenerateQRHandler(logger *zap.Logger) fiber.Handler {
    return func(c *fiber.Ctx) error {
        req := new(models.RequestPayload)
        if err := c.BodyParser(req); err != nil {
            logger.Error("Failed to parse request", zap.Error(err))
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Cannot parse JSON",
            })
        }

        qrCodeBase64, err := qr.GenerateQRCodeBase64(req.Payload)
        if err != nil {
            logger.Error("Failed to generate QR code", zap.Error(err))
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to generate QR code",
            })
        }

        logger.Info("QR code generated successfully")
        return c.JSON(models.ResponsePayload{QRCodeBase64: qrCodeBase64})
    }
}

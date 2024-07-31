package main

import (
	"github.com/Sagn1k/pixelQR/internal/handlers"
    "github.com/Sagn1k/pixelQR/internal/log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "go.uber.org/zap"
    "runtime/debug"
)

func main() {
    app := fiber.New()
    
    logger := log.NewLogger()

    app.Use(recover.New(recover.Config{
        EnableStackTrace: true,
        StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
            // Log the stack trace along with the error
            logger.Error("Panic occurred", zap.Any("error", e), zap.String("stack", string(debug.Stack())))
        },
    }))

    app.Post("/generate-qr", handlers.GenerateQRHandler(logger))

    logger.Info("Server is running on port 3000")
    app.Listen(":3000")
}

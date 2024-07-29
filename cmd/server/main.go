package main

import (
    "github.com/Sagn1k/pixelQR/internal/handlers"
    "github.com/Sagn1k/pixelQR/internal/log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    
    logger := log.NewLogger()

    app.Post("/generate-qr", handlers.GenerateQRHandler(logger))

    logger.Info("Server is running on port 3000")
    app.Listen(":3000")
}

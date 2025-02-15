package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var app *fiber.App = fiber.New()

func Serve() {
	log()
	routes()
	app.Listen(":3000")
}

func routes() {
	app.Post("/api/input", func(Ctx *fiber.Ctx) error {

		type WIFI struct {
			SSID     string `json:"ssid"`
			Password string `json:"password"`
		}

		var wifi []WIFI
		if err := Ctx.BodyParser(&wifi); err != nil {
			return Ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		for _, wifi := range wifi {
			fmt.Printf("Received WIFI SSID: %s, Password: %s\n", wifi.SSID, wifi.Password)
		}

		return Ctx.SendStatus(200)
	})
}

func log() {
	file, _ := os.OpenFile("logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output: file,
	}))
}

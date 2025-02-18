package server

import (
	"encoding/json"
	"os"
	"path/filepath"

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

		var dirPath string = "server/credentials"
		var filePath string = filepath.Join(dirPath, "creds.json")

		// Ensure the directory exists
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return Ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return Ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer file.Close()

		var encoder *json.Encoder = json.NewEncoder(file)
		if err := encoder.Encode(wifi); err != nil {
			return Ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return Ctx.SendStatus(200)
	})
}

func log() {
	app.Use(logger.New())
	file, _ := os.OpenFile("logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output: file,
	}))
}

package main

import (
	"github.com/anikets08/vidsume/controllers"
	"github.com/anikets08/vidsume/database"
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/health", controllers.Health)
}

func main() {
	app := fiber.New()
	database.ConnectDB()
	Routers(app)
	app.Listen(":3000")
}

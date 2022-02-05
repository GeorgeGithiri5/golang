package main

import (
	"example.com/go-auth/database"
	"example.com/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}

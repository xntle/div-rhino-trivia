package main

import (
	"github.com/gofiber/fiber/v2"
	databse "github.com/xntle/div-rhino-trivia/database"
)

func main() {
	databse.ConnectDb()
	app := fiber.New()


	app.Listen(":3000")
}
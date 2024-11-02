package main

import (
	"github.com/gofibar/fiber/v2"
	"github.com/xntle/div-rhino-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
} 
 
package main

import (
	"encode/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	if err := godotenv.Load(); err != nil {
		log.Println("Error .env")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "POST",
	}))

	routes.SetRoutes(app)

	app.Listen(":8080")
}

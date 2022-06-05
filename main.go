package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hberkayozdemir/hibemi-be/internal/user"
	"log"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	userRepository := user.NewRepository("mongodb://localhost:27017")

	userService := user.NewService(userRepository)

	userHandler := user.NewHandler(userService)

	userHandler.SetupApp(app)

	log.Fatal(app.Listen(":8080"))

}

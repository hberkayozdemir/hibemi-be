package main

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/hberkayozdemir/hibemi-be/internal/news"
	"github.com/hberkayozdemir/hibemi-be/internal/user"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	const DB_URL = "mongodb+srv://hbo_admin:pDXudWBy76rnVngE@hibemi.gtlntks.mongodb.net/?retryWrites=true&w=majority"
	var (
		apiKey    = "cQZU8AAqYsOBF6gV5YdFkZUulm0ce3dTqSQsG7IQmg3CzFMq3Ab9oRqpFOTtS6vF"
		secretKey = "zgthBiYZ8wAP1cTOQUyuMAVUjY7qqrwNRHGYL65ttBGMhQ8CoROjU76gYoFtQUOX"
	)
	binanceclient := binance.NewClient(apiKey, secretKey)
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {})
	c.Start()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	userRepository := user.NewRepository(DB_URL)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
	userHandler.SetupApp(app)
	newsRepository := news.NewRepository(DB_URL)
	newsService := news.NewService(newsRepository)
	newsHandler := news.NewHandler(newsService)
	newsHandler.SetupApp(app)

	prices, err := binanceclient.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}
	log.Fatal(app.Listen(":8080"))
}

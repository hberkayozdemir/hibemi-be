package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hberkayozdemir/hibemi-be/internal/client/coin_service"
	"github.com/hberkayozdemir/hibemi-be/internal/coin"
	"github.com/hberkayozdemir/hibemi-be/internal/news"
	"log"

	"github.com/hberkayozdemir/hibemi-be/internal/user"
	"github.com/robfig/cron/v3"
)

func main() {
	const DB_URL = "mongodb+srv://hbo_admin:pDXudWBy76rnVngE@hibemi.gtlntks.mongodb.net/?retryWrites=true&w=majority"

	client := coin_service.NewCoinGeckoClient(coin.Repository{})
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { client.FetchCoins() })
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

	log.Fatal(app.Listen(":8080"))

}

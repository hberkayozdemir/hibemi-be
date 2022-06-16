package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hberkayozdemir/hibemi-be/internal/banner"
	"github.com/hberkayozdemir/hibemi-be/internal/binance_spot"
	"github.com/hberkayozdemir/hibemi-be/internal/coin"
	"github.com/hberkayozdemir/hibemi-be/internal/favlist"
	"github.com/hberkayozdemir/hibemi-be/internal/news"
	"github.com/hberkayozdemir/hibemi-be/internal/transactions"
	"github.com/hberkayozdemir/hibemi-be/internal/user"
	"github.com/robfig/cron/v3"
)

var (
	apiKey    = "cQZU8AAqYsOBF6gV5YdFkZUulm0ce3dTqSQsG7IQmg3CzFMq3Ab9oRqpFOTtS6vF"
	secretKey = "zgthBiYZ8wAP1cTOQUyuMAVUjY7qqrwNRHGYL65ttBGMhQ8CoROjU76gYoFtQUOX"
)

var coins = []string{"ethereum", "bitcoin", "avalanche", "solona", "ripple", "near", "polkadot", "cordano", "maker", "chainlink", "stellar", "loopring", "dogecoin", "tron", "reef", "ontology", "bittorent", "algorand", "tether", "litecoin"}

func main() {

	const DB_URL = "mongodb+srv://hbo:Hbo.1998@hibemibe.zbozrlc.mongodb.net/?retryWrites=true&w=majority"

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
	binanceSpotRepository := binance_spot.NewRepository(DB_URL)
	binanceSpotService := binance_spot.NewService(binanceSpotRepository)
	binanceSpotHandler := binance_spot.NewHandler(binanceSpotService)
	binanceSpotHandler.SetupApp(app)
	coinRepository := coin.NewRepository(DB_URL)
	coinService := coin.NewService(coinRepository)
	coinHandler := coin.NewHandler(coinService)
	coinHandler.SetupApp(app)
	transactionRepository := transactions.NewRepository(DB_URL)
	transactionService := transactions.NewService(transactionRepository)
	transactionHandler := transactions.NewHandler(transactionService)
	transactionHandler.SetupApp(app)

	favlistRepository := favlist.NewRepository(DB_URL)
	favlistService := favlist.NewService(favlistRepository)
	favlistHandler := favlist.NewHandler(favlistService)
	favlistHandler.SetupApp(app)

	bannerRepository := banner.NewRepository(DB_URL)
	bannerService := banner.NewService(bannerRepository)
	bannerHandler := banner.NewHandler(bannerService)
	bannerHandler.SetupApp(app)

	port := SetPort()

	c := cron.New()
	c.AddFunc("@every 20s", func() {
		binanceSpotService.Repository.GetSpotsIteratable()

	})

	c.Start()
	log.Fatal(app.Listen(":" + port))
}

func SetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "8080"
	}

	return port
}

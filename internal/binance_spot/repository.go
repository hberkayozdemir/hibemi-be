package binance_spot

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Repository struct {
	MongoClient *mongo.Client
}

func NewRepository(uri string) Repository {
	env := os.Getenv("APP_ENV")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	if env == "test" {
		client.Database("ventures").Collection("spots").Drop(ctx)
	}

	return Repository{client}
}

func (r *Repository) UpdateDb(symbol []*binance.SymbolPrice) {
	collection := r.MongoClient.Database("ventures").Collection("spots")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for _, p := range symbol {
		if strings.Contains(p.Symbol, "USDT") {
			symbolEntity := convertSymbolPriceToSymbolPriceEntity(p)
			_, err := collection.InsertOne(ctx, &symbolEntity)
			fmt.Print(err == nil, p.Symbol)
		}

	}
}

func (r *Repository) GetSpotsIteratable() error {
	collection := r.MongoClient.Database("ventures").Collection("spots")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection.Drop(ctx)
	binanceClient := binance.NewClient(apiKey, secretKey)
	prices, err := binanceClient.NewListPricesService().Do(context.Background(), binance.WithHeaders(http.Header{
		"apiKey":    {apiKey},
		"secretKey": {secretKey},
		"symbols":   {"BNBUSDT", "ETHUSDT", "BTCUSDT", "XRPUSDT", "NEOUSDT", "DOTUSDT", "SOLUSDT", "AVAXUSDT", "WAVESUSDT", "LINKUSDT", "NEARUSDT"},
	}))
	r.UpdateDb(prices)
	return err
}

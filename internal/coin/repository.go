package coin

import (
	"context"
	"github.com/hberkayozdemir/hibemi-be/internal/coin_gecko"
	"go.mongodb.org/mongo-driver/bson"
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
		client.Database("ventures").Collection("users").Drop(ctx)
	}

	return Repository{client}
}

func (r *Repository) getAllSpots(page, size int) ([]Coins, int, error) {
	collection := r.MongoClient.Database("ventures").Collection("spots")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	options := options.Find()
	if size != 0 {
		options.SetSkip(int64(page * size))
		options.SetLimit(int64(size))
	}
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, nil
	}
	var coins []Coins
	for cur.Next(ctx) {
		userEntity := Coins{}
		err := cur.Decode(&userEntity)
		if err != nil {
			return nil, 0, nil
		}
		coins = append(coins, userEntity)
	}

	totalElements, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, nil
	}

	return coins, int(totalElements), nil
}

func (r *Repository) AddCoin(coin coin_gecko.CoinGeckoResponse) (map[string]interface{}, error) {
	collection := r.MongoClient.Database("ventures").Collection("coins")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, coin)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *Repository) GetAllCoins() ([]coin_gecko.CoinGeckoResponse, error) {
	collection := r.MongoClient.Database("ventures").Collection("coins")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var coins []coin_gecko.CoinGeckoResponse
	for cur.Next(ctx) {
		coingeckoEntity := coin_gecko.CoinGeckoResponseEntity{}
		err := cur.Decode(coingeckoEntity)
		if err != nil {
			return nil, err
		}
		coins = append(coins, CoinGeckoEntityToModel(coingeckoEntity))
	}

	return coins, nil
}

func (r *Repository) GetSpotsByIDList(idList []string) ([]coin_gecko.CoinGeckoResponse, error) {
	collection := r.MongoClient.Database("ventures").Collection("coins")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"symbol": bson.M{"$in": idList}}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var coins []coin_gecko.CoinGeckoResponse
	for cur.Next(ctx) {
		coingeckoEntity := coin_gecko.CoinGeckoResponseEntity{}
		err := cur.Decode(coingeckoEntity)
		if err != nil {
			return nil, err
		}
		coins = append(coins, CoinGeckoEntityToModel(coingeckoEntity))
	}

	return coins, nil
}

func (r *Repository) UpdateGeckoPrice(lowerSymbol string, price string) (*coin_gecko.CoinGeckoResponse, error) {
	collection := r.MongoClient.Database("ventures").Collection("coins")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"symbol": lowerSymbol}
	update := bson.D{{"$set",
		bson.D{
			{"price", price},
		},
	}}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

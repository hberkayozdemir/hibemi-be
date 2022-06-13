package coin

import (
	"context"
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

func (r *Repository) AddCoin(coin map[string]interface{}) (map[string]interface{}, error) {
	collection := r.MongoClient.Database("ventures").Collection("coins")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, coin)

	if err != nil {
		return nil, err
	}

	return r.GetNewsByID(news.ID)
}

package banner

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Repository struct {
	MongoClient *mongo.Client
}

func NewRepository(uri string) Repository {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	return Repository{client}
}

func (r *Repository) getBannerList() ([]Banner, error) {
	collection := r.MongoClient.Database("ventures").Collection("banner")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	options := options.Find()

	filter := bson.M{}
	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, nil
	}
	var banners []Banner
	for cur.Next(ctx) {
		bannerEntity := Banner{}
		err := cur.Decode(&bannerEntity)
		if err != nil {
			return nil, nil
		}
		banners = append(banners, bannerEntity)
	}
	return banners, nil
}

func (r *Repository) CreateBanner(banner Banner) (*Banner, error) {
	collection := r.MongoClient.Database("ventures").Collection("banner")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, banner)
	if err != nil {
		return nil, err
	}
	return r.GetBannerByID(banner.ID)
}

func (r *Repository) GetBannerByID(id string) (*Banner, error) {
	collection := r.MongoClient.Database("ventures").Collection("banner")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	cur := collection.FindOne(ctx, filter)
	if cur.Err() != nil {
		return nil, cur.Err()
	}

	if cur == nil {
		return nil, cur.Err()
	}

	bannerEntity := Banner{}
	err := cur.Decode(&bannerEntity)
	if err != nil {
		return nil, err
	}
	return &bannerEntity, nil
}

func (r *Repository) DeleteBanner(id string) error {
	collection := r.MongoClient.Database("ventures").Collection("banner")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

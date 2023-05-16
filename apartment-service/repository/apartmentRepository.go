package repository

import (
	"BookingPlatform/apartment-service/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type ApartmentRepository struct {
	Cli    *mongo.Client
	Logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*ApartmentRepository, error) {
	dburi := os.Getenv("MONGODB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &ApartmentRepository{
		Cli:    client,
		Logger: logger,
	}, nil
}
func (a *ApartmentRepository) Disconnect(ctx context.Context) error {
	err := a.Cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (a *ApartmentRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := a.Cli.Ping(ctx, readpref.Primary())
	if err != nil {
		a.Logger.Println(err)
	}

	dbs, err := a.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		a.Logger.Println(err)
	}
	fmt.Println(dbs)
}
func (a *ApartmentRepository) GetCollection() *mongo.Collection {
	bookingDatabase := a.Cli.Database("booking")
	apartmentsCollection := bookingDatabase.Collection("apartments")
	return apartmentsCollection
}

func (a *ApartmentRepository) Insert(apartment *model.Apartment) (*model.Apartment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	apartmentsCollection := a.GetCollection()

	result, err := apartmentsCollection.InsertOne(ctx, &apartment)
	if err != nil {
		a.Logger.Println(err)
		return nil, err
	}
	a.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return apartment, nil
}

func (a *ApartmentRepository) GetAll() (model.Apartments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	apartmentsCollection := a.GetCollection()

	var apartments model.Apartments
	apartmentsCursor, err := apartmentsCollection.Find(ctx, bson.M{})
	if err != nil {
		a.Logger.Println(err)
		return nil, err
	}
	if err = apartmentsCursor.All(ctx, &apartments); err != nil {
		a.Logger.Println(err)
		return nil, err
	}
	return apartments, nil
}

package repository

import (
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

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

type ReservationRepository struct {
	Cli    *mongo.Client
	Logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*ReservationRepository, error) {
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

	return &ReservationRepository{
		Cli:    client,
		Logger: logger,
	}, nil
}
func (r *ReservationRepository) Disconnect(ctx context.Context) error {
	err := r.Cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (r *ReservationRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Cli.Ping(ctx, readpref.Primary())
	if err != nil {
		r.Logger.Println(err)
	}

	dbs, err := r.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		r.Logger.Println(err)
	}
	fmt.Println(dbs)
}

package repository

import (
	"BookingPlatform/reservation-service/model"
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
	if dburi == "" {
		dburi = os.Getenv("mongodb+srv://draga:draga@cluster0.dlhjqkp.mongodb.net/?retryWrites=true&w=majority")
	}

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

func (fr *ReservationRepository) GetCollection() *mongo.Collection {
	bookingDatabase := fr.Cli.Database("booking")
	reservationsCollection := bookingDatabase.Collection("reservations")
	return reservationsCollection
}

func (rr *ReservationRepository) GetAll(guestId string) (model.Reservations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := rr.GetCollection()

	var reservations model.Reservations
	flightsCursor, err := reservationsCollection.Find(ctx, bson.M{"guestId": guestId})
	if err != nil {
		rr.Logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &reservations); err != nil {
		rr.Logger.Println(err)
		return nil, err
	}
	return reservations, nil
}

func (rr *ReservationRepository) Insert(reservation *model.Reservation) (*model.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationsCollection := rr.GetCollection()

	result, err := reservationsCollection.InsertOne(ctx, &reservation)
	if err != nil {
		rr.Logger.Println(err)
		return nil, err
	}
	rr.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return reservation, nil
}

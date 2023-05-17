package repository

import (
	"BookingPlatform/reservation-service/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (fr *ReservationRepository) GetCollectionRequests() *mongo.Collection {
	bookingDatabase := fr.Cli.Database("booking")
	reservationsCollection := bookingDatabase.Collection("reservations_requests")
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

func (rr *ReservationRepository) AcceptRequest(reservationRequest *model.ReservationRequset) (*model.ReservationRequset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationsCollection := rr.GetCollectionRequests()
	
	//delete all reservations with same startDate and endDate
	_, err2 := reservationsCollection.DeleteMany(ctx, bson.M{
		"startDate": reservationRequest.StartDate,
		"endDate":   reservationRequest.EndDate})
	if err2 != nil {
		rr.Logger.Println(err2)
		return nil, err2
	}

	result, err := reservationsCollection.InsertOne(ctx, &reservationRequest)
	if err != nil {
		rr.Logger.Println(err)
		return nil, err
	}

	rr.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return reservationRequest, nil
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

func (rr *ReservationRepository) InsertReservationRequest(reservation_request *model.ReservationRequset) (*model.ReservationRequset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationsCollection := rr.GetCollection()
	reservationsRequestsCollection := rr.GetCollectionRequests()

	var reservations model.Reservations
	reservationsCursor, err := reservationsCollection.Find(ctx, bson.M{
		"apartmentId": reservation_request.ApartmentID})
	if err != nil {
		rr.Logger.Println(err)
		return nil, err
	}
	if err = reservationsCursor.All(ctx, &reservations); err != nil {
		rr.Logger.Println(err)
		return nil, err
	}

	fmt.Println("*****************************************  ", len(reservations))
	if len(reservations) < 1 {
		return nil, err
	}

	req_startDate, _ := time.Parse("02-01-2006", reservation_request.StartDate)
	req_endDate, _ := time.Parse("02-01-2006", reservation_request.EndDate)
	for i := 0; i < len(reservations); i++ {
		res_startDate, _ := time.Parse("02-01-2006", reservations[i].StartDate)
		res_endDate, _ := time.Parse("02-01-2006", reservations[i].EndDate)
		if req_startDate.After(res_startDate) && req_startDate.Before(res_endDate) {
			return nil, err
		}
		if req_endDate.After(res_startDate) && req_endDate.Before(res_endDate) {
			return nil, err
		}
	}

	result, err := reservationsRequestsCollection.InsertOne(ctx, &reservation_request)
	if err != nil {
		rr.Logger.Println(err)
		return nil, err
	}
	rr.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return reservation_request, nil
}

func (rr *ReservationRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := rr.GetCollection()

	var reservation model.Reservation
	objID, _ := primitive.ObjectIDFromHex(id)
	err := reservationsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&reservation)
	if err != nil {
		rr.Logger.Println(err)
		return err
	}

	today := time.Now()
	res_time, _ := time.Parse("02-01-2006", reservation.StartDate)
	tomorrow := today.Add(24 * time.Hour)
	if tomorrow.After(res_time) {
		return nil
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := reservationsCollection.DeleteOne(ctx, filter)

	if err != nil {
		rr.Logger.Println(err)
		return err
	}

	rr.Logger.Printf("Documents deleted : %v\n", result.DeletedCount)
	return nil
}

func (rr *ReservationRepository) DeleteRequest(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsRequestsCollection := rr.GetCollectionRequests()
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := reservationsRequestsCollection.DeleteOne(ctx, filter)
	if err != nil {
		rr.Logger.Println(err)
		return err
	}
	rr.Logger.Printf("Documents deleted: %v\n", result.DeletedCount)

	return nil
}

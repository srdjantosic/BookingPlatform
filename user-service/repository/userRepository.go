package repository

import (
	"BookingPlatform/user-service/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserRepository struct {
	Cli    *mongo.Client
	Logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*UserRepository, error) {
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
	return &UserRepository{
		Cli:    client,
		Logger: logger,
	}, nil
}
func (u *UserRepository) Disconnect(ctx context.Context) error {
	err := u.Cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *UserRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.Cli.Ping(ctx, readpref.Primary())
	if err != nil {
		u.Logger.Println(err)
	}

	dbs, err := u.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		u.Logger.Println(err)
	}
	fmt.Println(dbs)
}

func (ur *UserRepository) GetCollection() *mongo.Collection {
	bookingDatabase := ur.Cli.Database("booking")
	usersCollection := bookingDatabase.Collection("users")
	return usersCollection
}

func (ur *UserRepository) GetCollectionApartments() *mongo.Collection {
	bookingDatabase := ur.Cli.Database("booking")
	apartmentsCollection := bookingDatabase.Collection("apartments")
	return apartmentsCollection
}

func (ur *UserRepository) GetCollectionReservations() *mongo.Collection {
	bookingDatabase := ur.Cli.Database("booking")
	reservationsCollection := bookingDatabase.Collection("reservations")
	return reservationsCollection
}

func (ur *UserRepository) GetCollectionReservationsRequests() *mongo.Collection {
	bookingDatabase := ur.Cli.Database("booking")
	reservationsRequestsCollection := bookingDatabase.Collection("reservations_requests")
	return reservationsRequestsCollection
}

func (ur *UserRepository) GetOne(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.GetCollection()

	var user model.User
	objId, _ := primitive.ObjectIDFromHex(id)
	err := usersCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Insert(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := ur.GetCollection()

	result, err := usersCollection.InsertOne(ctx, &user)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	ur.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return user, nil
}
func (ur *UserRepository) FindByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.GetCollection()

	var user model.User
	err := usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	return &user, err
}

func (ur *UserRepository) FindByUsernameAndPassword(username string, password string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.GetCollection()

	var user model.User
	err := usersCollection.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetApartmentById(apartmentId primitive.ObjectID) (*model.Apartment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ur.Logger.Println("^^^^^^^^^^^  ", apartmentId, " ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	apartmentssCollection := ur.GetCollectionApartments()

	var apartment model.Apartment
	err := apartmentssCollection.FindOne(ctx, bson.M{"_id": apartmentId}).Decode(&apartment)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}

	return &apartment, nil
}

func (ur *UserRepository) Update(id string, user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.GetCollection()

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"password":  user.Password,
		"username":  user.Username,
		"address":   user.Address,
	}}

	result, err := usersCollection.UpdateOne(ctx, filter, update)
	ur.Logger.Printf("Documents matched: %v\n", result.MatchedCount)
	ur.Logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		ur.Logger.Println(err)
		return err
	}
	return nil
}

func (ur *UserRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.GetCollection()

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}
	result, err := usersCollection.DeleteOne(ctx, filter)
	if err != nil {
		ur.Logger.Println(err)
		return err
	}
	ur.Logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (ur *UserRepository) DeleteReservation(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := ur.GetCollectionReservations()

	var reservation model.Reservation
	objID, _ := primitive.ObjectIDFromHex(id)
	err := reservationsCollection.FindOne(ctx, bson.M{"id": objID}).Decode(&reservation)
	if err != nil {
		ur.Logger.Println(err)
		return err
	}

	today := time.Now()
	res_time, _ := time.Parse("02-01-2006", reservation.StartDate)
	tomorrow := today.Add(24 * time.Hour)
	if tomorrow.After(res_time) {
		return nil
	}

	filter := bson.D{{Key: "id", Value: objID}}
	result, err := reservationsCollection.DeleteOne(ctx, filter)

	if err != nil {
		ur.Logger.Println(err)
		return err
	}

	ur.Logger.Printf("Documents deleted : %v\n", result.DeletedCount)
	return nil
}

func (ur *UserRepository) DeleteRequest(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := ur.GetCollectionReservationsRequests()

	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := reservationsCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		ur.Logger.Println(err)
		return err
	}

	ur.Logger.Printf("Documents deleted : %v\n", result.DeletedCount)
	return nil
}

func (ur *UserRepository) GetAllReservationsByUser(guestId string) (model.Reservations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := ur.GetCollectionReservations()

	objID, _ := primitive.ObjectIDFromHex(guestId)
	fmt.Println("**************************", objID, " *************************************")
	filter := bson.D{{Key: "guestId", Value: objID}}

	var reservations model.Reservations
	//flightsCursor, err := reservationsCollection.Find(ctx, bson.M{"guestId": objID})
	flightsCursor, err := reservationsCollection.Find(ctx, filter)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &reservations); err != nil {
		ur.Logger.Println(err)
		ur.Logger.Println("*************************************************")
		return nil, err
	}
	return reservations, nil
}

func (ur *UserRepository) GetAllReservationsByApartment(apartmentId primitive.ObjectID) (model.Reservations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := ur.GetCollectionReservations()

	fmt.Println("**************************", apartmentId, " *************************************")
	filter := bson.D{{Key: "apartmentId", Value: apartmentId}}

	var reservations model.Reservations
	//flightsCursor, err := reservationsCollection.Find(ctx, bson.M{"guestId": objID})
	flightsCursor, err := reservationsCollection.Find(ctx, filter)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &reservations); err != nil {
		ur.Logger.Println(err)
		ur.Logger.Println("*************************************************")
		return nil, err
	}
	return reservations, nil
}

func (ur *UserRepository) InsertReservation(reservation *model.Reservation) (*model.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationsCollection := ur.GetCollectionReservations()

	result, err := reservationsCollection.InsertOne(ctx, &reservation)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	ur.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return reservation, nil
}

func (ur *UserRepository) FindAllApartmentsByHostId(hostId primitive.ObjectID) (model.Apartments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	apartmentsCollection := ur.GetCollectionApartments()

	filter := bson.D{{Key: "hostId", Value: hostId}}

	var apartments model.Apartments
	flightsCursor, err := apartmentsCollection.Find(ctx, filter)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &apartments); err != nil {
		ur.Logger.Println(err)
		ur.Logger.Println("*************************************************")
		return nil, err
	}
	return apartments, nil
}

/*
// TODO mislim da je u for petlji problem jer ne uspe da doda zahtev, nadjeni apartmani po host id su dobri

	func (ur *UserRepository) FindAllReservationRequestsByApartments(apartments model.Apartments) (model.ReservationRequests, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		reservationRequestsCollection := ur.GetCollectionReservationsRequests()

		var reservationRequests model.ReservationRequests

		for i := 0; i < len(apartments); i++ {
			fmt.Println("apartmentId: " + apartments[i].ID.String())
			reservations, err := reservationRequestsCollection.Find(ctx, bson.M{"apartmentId": apartments[i].ID})
			if err != nil {
				ur.Logger.Println(err)
				return nil, err
			}
			reservations.All(ctx, &reservationRequests)
			fmt.Println("DUZINAAAAAAAAAAAAAAAAAAAAA")
			fmt.Println(len(reservationRequests))

			reservations.Decode(reservationRequests)
		}

		return reservationRequests, nil
	}
*/
func (ur *UserRepository) GetAllReservationRequestsByApartment(apartmentId primitive.ObjectID) (model.ReservationRequests, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationRequestsCollection := ur.GetCollectionReservationsRequests()

	fmt.Println("**************************", apartmentId, " *************************************")
	filter := bson.D{{Key: "apartmentId", Value: apartmentId}}

	var reservations model.ReservationRequests
	//flightsCursor, err := reservationsCollection.Find(ctx, bson.M{"guestId": objID})
	flightsCursor, err := reservationRequestsCollection.Find(ctx, filter)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &reservations); err != nil {
		ur.Logger.Println(err)
		ur.Logger.Println("*************************************************")
		return nil, err
	}
	return reservations, nil
}

func (ur *UserRepository) InsertReservationRequest(reservationRequest *model.ReservationRequset) (*model.ReservationRequset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationsCollection := ur.GetCollectionReservationsRequests()

	result, err := reservationsCollection.InsertOne(ctx, &reservationRequest)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}
	ur.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return reservationRequest, nil
}

func (ur *UserRepository) AcceptRequest(id string) (*model.ReservationRequset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	requestsCollection := ur.GetCollectionReservationsRequests()

	var reservationRequest model.ReservationRequset
	objID, _ := primitive.ObjectIDFromHex(id)
	err := requestsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&reservationRequest)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}

	//delete all requests with same startDate and endDate
	_, err2 := requestsCollection.DeleteMany(ctx, bson.M{
		"startDate": reservationRequest.StartDate,
		"endDate":   reservationRequest.EndDate})
	if err2 != nil {
		ur.Logger.Println(err2)
		return nil, err2
	}

	reservationCollection := ur.GetCollectionReservations()

	reservation := model.Reservation{
		ID:           reservationRequest.ID,
		GuestID:      reservationRequest.UserID,
		ApartmentID:  reservationRequest.ApartmentID,
		StartDate:    reservationRequest.StartDate,
		EndDate:      reservationRequest.EndDate,
		GuestsNumber: reservationRequest.GuestsNumber,
	}

	result, err := reservationCollection.InsertOne(ctx, &reservation)
	if err != nil {
		ur.Logger.Println(err)
		return nil, err
	}

	ur.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	return &reservationRequest, nil
}

func (ur *UserRepository) DeleteHostsApartments(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bookingDatabase := ur.Cli.Database("booking")
	apartmentCollection := bookingDatabase.Collection("apartments")

	filter := bson.D{{Key: "_id", Value: id}}
	result, err := apartmentCollection.DeleteOne(ctx, filter)
	if err != nil {
		ur.Logger.Println(err)
		return err
	}
	ur.Logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

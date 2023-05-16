package repository

import (
	"BookingPlatform/user-service/model"
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

type UserRepository struct {
	Cli    *mongo.Client
	Logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*UserRepository, error) {
	//dburi := os.Getenv("MONGODB_URI")
	dburi := os.Getenv("mongodb+srv://draga:draga@cluster0.dlhjqkp.mongodb.net/?retryWrites=true&w=majority")

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

func (ur *UserRepository) GetOne(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.GetCollection()

	var user model.User
	objId, _ := primitive.ObjectIDFromHex(id)
	err := usersCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
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

package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type ReservationRequset struct {
	ID           primitive.ObjectID `bson:"id" json:"id"`
	UserID       primitive.ObjectID `bson:"userId" json:"userId"`
	ApartmentID  primitive.ObjectID `bson:"apartmentId" json:"apartmentId"`
	StartDate    string             `bson:"startDate" json:"startDate"`
	EndDate      string             `bson:"endDate" json:"endDate"`
	GuestsNumber int                `bson:"guestsNumber" json:"guestsNumber"`
}

type ReservationRequests []*ReservationRequset

func (r *ReservationRequests) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(r)
}

func (r *ReservationRequset) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(r)
}
func (r *ReservationRequset) FromJSON(rd io.Reader) error {
	d := json.NewDecoder(rd)
	return d.Decode(r)
}

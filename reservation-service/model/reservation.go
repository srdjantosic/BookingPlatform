package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Reservation struct {
	ID           primitive.ObjectID `bson:"id" json:"id"`
	GuestID      primitive.ObjectID `bson:"guestId" json:"guestId"`
	ApartmentID  primitive.ObjectID `bson:"apartmentId" json:"apartmentId"`
	StartDate    string             `bson:"startDate" json:"startDate"`
	EndDate      string             `bson:"endDate" json:"endDate"`
	GuestsNumber int                `bson:"guestsNumber" json:"guestsNumber"`
}

type Reservations []*Reservation

func (r *Reservations) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(r)
}

func (r *Reservation) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(r)
}
func (r *Reservation) FromJSON(rd io.Reader) error {
	d := json.NewDecoder(rd)
	return d.Decode(r)
}

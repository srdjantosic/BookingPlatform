package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type User struct {
	ID                    primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	FirstName             string             `bson:"firstName" json:"firstName"`
	LastName              string             `bson:"lastName" json:"lastName"`
	Email                 string             `bson:"email" json:"email"`
	Password              string             `bson:"password" json:"password"`
	Username              string             `bson:"username" json:"username"`
	Address               string             `bson:"address" json:"address"`
	Role                  string             `bson:"role" json:"role"`
	DeclineRequestCounter int                `bson:"declineRequestCounter" json:"declineRequestCounter"`
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}
func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

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

type Apartment struct {
	ID                   primitive.ObjectID `bson:"_id" json:"id"`
	HostId               primitive.ObjectID `bson:"hostId" json:"hostId"`
	Name                 string             `bson:"name" json:"name"`
	Location             string             `bson:"location" json:"location"`
	Benefits             string             `bson:"benefits" json:"benefits"`
	MinGuestsNumber      int                `bson:"minGuestsNumber" json:"minGuestsNumber"`
	MaxGuestsNumber      int                `bson:"maxGuestsNumber" json:"maxGuestsNumber"`
	AutomaticReservation bool               `bson:"automaticReservation" json:"automaticReservation"`
	Pricelist            []*PricelistItem   `bson:"pricelist" json:"pricelist"`
}

type Apartments []*Apartment

func (a *Apartments) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *Apartment) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *Apartment) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

type Unit int

const (
	PricePerPerson Unit = iota
	PriceForApartmentUnit
)

type PricelistItem struct {
	AvailabilityStartDate string `bson:"availabilityStartDate" json:"availabilityStartDate"`
	AvailabilityEndDate   string `bson:"availabilityEndDate" json:"availabilityEndDate"`
	Price                 int    `bson:"price" json:"price"`
	UnitPrice             Unit   `bson:"unitPrice" json:"unitPrice"`
}

func (p *PricelistItem) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *PricelistItem) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type AvailableApartment struct {
	Apartment  *Apartment
	TotalPrice int
	UnitPrice  int
}

type AvailableApartments []*AvailableApartment

func (a *AvailableApartments) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *AvailableApartment) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *AvailableApartment) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

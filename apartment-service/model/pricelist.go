package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Unit int

const (
	PricePerPerson Unit = iota
	PriceForApartmentUnit
)

type Pricelist struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	ApartmentId    primitive.ObjectID `bson:"apartmentId" json:"apartmentId"`
	PricelistItems []PricelistItem    `bson:"pricelistItems" json:"pricelistItems"`
}

type PricelistItem struct {
	AvailabilityStartDate string `bson:"availabilityStartDate" json:"availabilityStartDate"`
	AvailabilityEndDate   string `bson:"availabilityEndDate" json:"availabilityEndDate"`
	Price                 int    `bson:"price" json:"price"`
	UnitPrice             Unit   `bson:"unitPrice" json:"unitPrice"`
}

func (p *Pricelist) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Pricelist) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *PricelistItem) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *PricelistItem) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

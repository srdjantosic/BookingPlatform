package model

import (
	"encoding/json"
	"io"
)

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

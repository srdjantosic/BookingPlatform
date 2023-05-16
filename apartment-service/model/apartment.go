package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Apartment struct {
	ID       primitive.ObjectID `bson:"id" json:"id"`
	HostId   primitive.ObjectID `bson:"hostId" json:"hostId"`
	Name     string             `bson:"name" json:"name"`
	Location string             `bson:"location" json:"location"`
	Benefits string             `bson:"benefits" json:"benefits"`
	//FOTOGRAFIJE???
	minGuestsNumber int `bson:"minGuestsNumber" json:"minGuestsNumber"`
	maxGuestsNumber int `bson:"maxGuestsNumber" json:"maxGuestsNumber"`
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

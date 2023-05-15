package model

import (
	"encoding/json"
	"io"
)

type Apartment struct {
	ID       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`
	Benefits string `bson:"benefits" json:"benefits"`
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

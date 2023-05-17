package model

import (
	"encoding/json"
	"io"
)

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

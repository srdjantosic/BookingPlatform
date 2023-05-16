package service

import (
	"BookingPlatform/apartment-service/model"
	"BookingPlatform/apartment-service/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type ApartmentService struct {
	Repo   *repository.ApartmentRepository
	Logger *log.Logger
}

func NewApartmentService(r *repository.ApartmentRepository, l *log.Logger) *ApartmentService {
	return &ApartmentService{r, l}
}

func (as *ApartmentService) Insert(apartment *model.Apartment, userRole string, userId string) (*model.Apartment, error) {
	if userRole == "HOST" {
		apartment.HostId, _ = primitive.ObjectIDFromHex(userId)
		return as.Repo.Insert(apartment)
	}
	return nil, fmt.Errorf("You are not authorized for this function!")
}

func (as *ApartmentService) GetAll() (model.Apartments, error) {
	apartments, err := as.Repo.GetAll()
	if err != nil {
		as.Logger.Println(err)
		return nil, err
	}
	return apartments, nil
}

package service

import (
	"BookingPlatform/apartment-service/repository"
	"log"
)

type ApartmentService struct {
	Repo   *repository.ApartmentRepository
	Logger *log.Logger
}

func NewApartmentService(r *repository.ApartmentRepository, l *log.Logger) *ApartmentService {
	return &ApartmentService{r, l}
}

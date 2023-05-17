package service

import (
	"BookingPlatform/reservation-service/repository"
	"log"
)

type ReservationService struct {
	Repo   *repository.ReservationRepository
	Logger *log.Logger
}

func NewReservationService(r *repository.ReservationRepository, l *log.Logger) *ReservationService {
	return &ReservationService{r, l}
}

func (rs *ReservationService) Delete(id string) error {
	return nil
}

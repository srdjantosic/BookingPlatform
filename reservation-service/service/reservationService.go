package service

import (
	"BookingPlatform/reservation-service/model"
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

func (rs *ReservationService) Insert(reservation *model.Reservation) (*model.Reservation, error) {
	return rs.Repo.Insert(reservation)
}

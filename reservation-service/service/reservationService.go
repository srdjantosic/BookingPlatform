package service

import (
	"BookingPlatform/reservation-service/model"
	"BookingPlatform/reservation-service/pb"
	"BookingPlatform/reservation-service/repository"
	"context"
	"log"
)

type ReservationService struct {
	Repo   *repository.ReservationRepository
	Logger *log.Logger
}

func NewReservationService(r *repository.ReservationRepository, l *log.Logger) *ReservationService {
	return &ReservationService{r, l}
}

type UserReservationService struct {
	pb.UnimplementedUserReservationServiceServer
}

func (rs *ReservationService) GetIfNoReservations(ctx context.Context, req *pb.GetReservationsForUserRequest) (*pb.GetReservationsForUserResponse, error) {
	reservations, err := rs.Repo.GetAll(req.GuestId)

	if err != nil {
		return nil, err
	}

	if len(reservations) == 0 {
		return &pb.GetReservationsForUserResponse{
			Message: true,
		}, nil
	} else {
		return &pb.GetReservationsForUserResponse{
			Message: false,
		}, nil
	}

}

func (rs *ReservationService) Insert(reservation *model.Reservation) (*model.Reservation, error) {
	return rs.Repo.Insert(reservation)
}

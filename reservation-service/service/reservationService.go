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

//type UserReservationService struct {
//	pb.UnimplementedUserReservationServiceServer
//}

func NewReservationService(r *repository.ReservationRepository, l *log.Logger) *ReservationService {
	return &ReservationService{r, l}
}

func (rs *ReservationService) Insert(reservation *model.Reservation) (*model.Reservation, error) {
	return rs.Repo.Insert(reservation)
}

func (rs *ReservationService) InsertReservationRequest(reservationRequest *model.ReservationRequset) (*model.ReservationRequset, error) {
	return rs.Repo.InsertReservationRequest(reservationRequest)
}

func (rs *ReservationService) Delete(id string) error {
	return rs.Repo.Delete(id)
}

//func (rs *ReservationService) GetReservationByGuestId(ctx context.Context, req *pb.GetReservationRequest) (*pb.GetReservationResponse, error) {
//	reservations, err := rs.Repo.GetAll(req.GuestId)
//
//	if err != nil {
//		return nil, err
//	}
//
//	if len(reservations) == 0 {
//		return &pb.GetReservationResponse{
//			Reservations: "true",
//		}, nil
//	} else {
//		return &pb.GetReservationResponse{
//			Reservations: false,
//		}, nil
//	}
//
//}

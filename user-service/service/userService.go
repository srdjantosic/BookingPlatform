package service

import (
	"BookingPlatform/user-service/model"
	"BookingPlatform/user-service/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserService struct {
	Repo   *repository.UserRepository
	Logger *log.Logger
}

func NewUserService(r *repository.UserRepository, l *log.Logger) *UserService {
	return &UserService{r, l}
}

func (us *UserService) Insert(user *model.User) (*model.User, error) {

	_, err := us.Repo.FindByUsername(user.Username)

	if err != nil {
		return us.Repo.Insert(user)
	}
	return nil, err
}

func (us *UserService) FindByUsernameAndPassword(username string, password string) (*model.User, error) {
	return us.Repo.FindByUsernameAndPassword(username, password)
}

func (us *UserService) Update(id string, userToUpdate *model.User) error {
	user, err := us.Repo.GetOne(id)

	if err == nil && (user.Role == "HOST" || user.Role == "GUEST") {
		if userToUpdate.FirstName != "" {
			user.FirstName = userToUpdate.FirstName
		}
		if userToUpdate.LastName != "" {
			user.LastName = userToUpdate.LastName
		}
		if userToUpdate.Email != "" {
			user.Email = userToUpdate.Email
		}
		if userToUpdate.Address != "" {
			user.Address = userToUpdate.Address
		}
		if userToUpdate.Password != "" {
			user.Password = userToUpdate.Password
		}
		if userToUpdate.Username != "" {
			user.Username = userToUpdate.Username
		}
		err = us.Repo.Update(id, user)
		if err != nil {
			return fmt.Errorf("Something went wrong!")
		}
		return nil
	}
	return fmt.Errorf("You are not authorized for this function!")
}
func (us *UserService) DeleteReservation(id string) error {
	return us.Repo.DeleteReservation(id)
}

func (us *UserService) DeleteRequest(id string) error {
	return us.Repo.DeleteRequest(id)
}

func (us *UserService) GetAllReservationsByUser(guestId string) (model.Reservations, error) {
	fmt.Println("USOOOOOOOOOOOOOOOOOOOOOO2222222222222222222")
	reservations, err := us.Repo.GetAllReservationsByUser(guestId)
	if err != nil {
		us.Logger.Println(err)
		return nil, err
	}
	return reservations, nil
}

func (us *UserService) InsertReservation(reservation *model.Reservation) (*model.Reservation, error) {
	return us.Repo.InsertReservation(reservation)
}

// TODO
func (us *UserService) FindAllApartmentsByHostId(hostId primitive.ObjectID) (model.Apartments, error) {

	apartments, err := us.Repo.FindAllApartmentsByHostId(hostId)
	if err != nil {
		us.Logger.Println(err)
		return nil, err
	}

	return apartments, nil
}

func (us *UserService) FindAllReservationRequestsByApartments(apartments model.Apartments) (model.ReservationRequests, error) {

	reservationRequests, err := us.Repo.FindAllReservationRequestsByApartments(apartments)
	if err != nil {
		us.Logger.Println(err)
		return nil, err
	}

	return reservationRequests, nil
}

func (us *UserService) InsertReservationRequest(reservation *model.ReservationRequset) (*model.ReservationRequset, error) {
	apartment, err := us.Repo.GetApartmentById(reservation.ApartmentID)
	if err != nil {
		us.Logger.Println(err)
		return nil, err
	}

	if apartment.AutomaticReservation == true {
		newReservation := new(model.Reservation)
		newReservation.StartDate = reservation.StartDate
		newReservation.ID = primitive.NewObjectID()
		newReservation.EndDate = reservation.EndDate
		newReservation.ApartmentID = reservation.ApartmentID
		newReservation.GuestID = reservation.UserID
		newReservation.GuestsNumber = reservation.GuestsNumber
		us.Repo.InsertReservation(newReservation)
		return reservation, err
	}

	us.Logger.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^FALSE^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	return us.Repo.InsertReservationRequest(reservation)
}

func (us *UserService) AcceptRequest(id string) (*model.ReservationRequset, error) {

	return us.Repo.AcceptRequest(id)
}
func (us *UserService) Delete(id string, role string) error {
	return nil
}

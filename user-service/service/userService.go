package service

import (
	"BookingPlatform/user-service/model"
	"BookingPlatform/user-service/repository"
	"fmt"
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

func (us *UserService) GetAllReservationsByUser(guestId string) (model.Reservations, error) {
	reservations, err := us.Repo.GetAllReservationsByUser(guestId)
	if err != nil {
		us.Logger.Println(err)
		return nil, err
	}
	return reservations, nil
}

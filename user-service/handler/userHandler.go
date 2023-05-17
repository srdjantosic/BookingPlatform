package handler

import (
	"BookingPlatform/user-service/model"
	"BookingPlatform/user-service/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type KeyProduct struct{}
type UserHandler struct {
	Service *service.UserService
	Logger  *log.Logger
}

func NewUserHandler(s *service.UserService, l *log.Logger) *UserHandler {
	return &UserHandler{s, l}
}

func (u *UserHandler) DatabaseName(ctx context.Context) {
	dbs, err := u.Service.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}
func (u *UserHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *UserHandler) MiddlewareReservationDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.Reservation{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *UserHandler) MiddlewareRequestDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.ReservationRequset{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
func (u *UserHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.Logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
func (u *UserHandler) Insert(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	user.ID = primitive.NewObjectID()

	createdUser, err := u.Service.Insert(user)

	if createdUser == nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) GetUserByUsernameAndPassword(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	username := vars["username"]
	password := vars["password"]

	user, err := u.Service.FindByUsernameAndPassword(username, password)

	if err != nil {
		fmt.Println("Error while logging in.")
		rw.WriteHeader(http.StatusBadRequest)
	}

	user.ToJSON(rw)
	rw.WriteHeader(http.StatusOK)

}

func (u *UserHandler) Update(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	user := h.Context().Value(KeyProduct{}).(*model.User)

	err := u.Service.Update(id, user)
	if err != nil {
		u.Logger.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) DeleteReservation(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	uh.Service.DeleteReservation(id)
	rw.WriteHeader(http.StatusNoContent)
}

func (uh *UserHandler) GetAllReservationsByUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	reservations, err := uh.Service.GetAllReservationsByUser(id)

	uh.Logger.Print("**************************", id, " *************************************")

	if err != nil {
		uh.Logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
	}
	if reservations == nil {
		return
	}
	err = reservations.ToJson(rw)
	if err != nil {
		uh.Logger.Print("Unable to convert to json :", err)
		rw.WriteHeader(http.StatusConflict)
	}
	rw.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) InsertReservation(rw http.ResponseWriter, h *http.Request) {
	reservation := h.Context().Value(KeyProduct{}).(*model.Reservation)
	reservation.ID = primitive.NewObjectID()
	fmt.Println(reservation.StartDate, " ", reservation.EndDate)

	createdReservation, err := uh.Service.InsertReservation(reservation)
	if createdReservation == nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) InsertReservationRequest(rw http.ResponseWriter, h *http.Request) {
	reservationRequest := h.Context().Value(KeyProduct{}).(*model.ReservationRequset)
	reservationRequest.ID = primitive.NewObjectID()

	createdReservation, err := uh.Service.InsertReservationRequest(reservationRequest)
	if createdReservation == nil {
		uh.Logger.Printf("*******************************************************")
		rw.WriteHeader(http.StatusBadRequest)
	}
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}

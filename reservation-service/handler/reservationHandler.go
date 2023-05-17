package handler

import (
	"BookingPlatform/reservation-service/model"
	"BookingPlatform/reservation-service/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type KeyProduct struct{}
type ReservationHandler struct {
	Service *service.ReservationService
	Logger  *log.Logger
}

func NewReservationHandler(s *service.ReservationService, l *log.Logger) *ReservationHandler {
	return &ReservationHandler{s, l}
}
func (r *ReservationHandler) DatabaseName(ctx context.Context) {
	dbs, err := r.Service.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}

func (r *ReservationHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		reservation := &model.Reservation{}
		err := reservation.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			r.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, reservation)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (r *ReservationHandler) MiddlewareRequestDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		reservationRequest := &model.ReservationRequset{}
		err := reservationRequest.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			r.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, reservationRequest)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
func (r *ReservationHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		r.Logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (rh *ReservationHandler) Insert(rw http.ResponseWriter, h *http.Request) {
	reservation := h.Context().Value(KeyProduct{}).(*model.Reservation)
	reservation.ID = primitive.NewObjectID()

	createdReservation, err := rh.Service.Insert(reservation)
	if createdReservation == nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}

func (rh *ReservationHandler) InsertReservationRequest(rw http.ResponseWriter, h *http.Request) {
	reservation_request := h.Context().Value(KeyProduct{}).(*model.ReservationRequset)
	reservation_request.ID = primitive.NewObjectID()

	/*ako je za apartman automatska potvrda
	model.Reservation reservation := model.Reservation{
		ID:           reservation_request.ID,
		GuestID:      reservation_request.UserID,
		ApartmentID:  reservation_request.ApartmentID,
		StartDate:    reservation_request.StartDate,
		EndDate:      reservation_request.EndDate,
		GuestsNumber: reservation_request.GuestsNumber,
	}

	createdReservation, err := rh.Service.Insert(&reservation)
	*/

	createdReservationRequest, err := rh.Service.InsertReservationRequest(reservation_request)
	if createdReservationRequest == nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}

func (rh *ReservationHandler) DeleteReservation(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	rh.Service.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

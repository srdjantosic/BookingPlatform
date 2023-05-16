package handler

import (
	"BookingPlatform/apartment-service/model"
	"BookingPlatform/apartment-service/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type KeyProduct struct{}
type ApartmentHandler struct {
	Service *service.ApartmentService
	Logger  *log.Logger
}

func NewApartmentHandler(s *service.ApartmentService, l *log.Logger) *ApartmentHandler {
	return &ApartmentHandler{s, l}
}
func (a *ApartmentHandler) DatabaseName(ctx context.Context) {
	dbs, err := a.Service.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}

func (a *ApartmentHandler) MiddlewareApartmentDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		apartment := &model.Apartment{}
		err := apartment.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			a.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, apartment)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
func (a *ApartmentHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		a.Logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (a *ApartmentHandler) Insert(rw http.ResponseWriter, h *http.Request) {
	apartment := h.Context().Value(KeyProduct{}).(*model.Apartment)
	apartment.ID = primitive.NewObjectID()
	vars := mux.Vars(h)
	userRole := vars["userRole"]
	userId := vars["userId"]

	createdApartment, err := a.Service.Insert(apartment, userRole, userId)
	if createdApartment == nil || err != nil {
		a.Logger.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusCreated)
}

func (a *ApartmentHandler) GetAll(rw http.ResponseWriter, h *http.Request) {
	apartments, err := a.Service.GetAll()
	if err != nil {
		a.Logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusBadRequest)
	}
	if apartments == nil {
		return
	}
	err = apartments.ToJson(rw)
	if err != nil {
		a.Logger.Print("Unable to convert to json :", err)
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusOK)
}

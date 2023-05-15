package handler

import (
	"BookingPlatform/apartment-service/model"
	"BookingPlatform/apartment-service/service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

func (a *ApartmentHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
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

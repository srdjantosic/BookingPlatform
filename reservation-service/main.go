package main

import (
	"BookingPlatform/reservation-service/handler"
	"BookingPlatform/reservation-service/repository"
	"BookingPlatform/reservation-service/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
)

func main() {

	fmt.Println("RESERVATION")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	reservationLogger := log.New(os.Stdout, "[reservation-store] ", log.LstdFlags)

	reservationRepository, err := repository.New(timeoutContext, reservationLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer reservationRepository.Disconnect(timeoutContext)

	reservationRepository.Ping()

	reservationService := service.NewReservationService(reservationRepository, logger)
	reservationHandler := handler.NewReservationHandler(reservationService, logger)

	reservationHandler.DatabaseName(timeoutContext)

	router := mux.NewRouter()
	router.Use(reservationHandler.MiddlewareContentTypeSet)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/insert", reservationHandler.Insert)
	postRouter.Use(reservationHandler.MiddlewareUserDeserialization)

	postRequestRouter := router.Methods(http.MethodPost).Subrouter()
	postRequestRouter.HandleFunc("/insertRequest", reservationHandler.InsertReservationRequest)
	postRequestRouter.Use(reservationHandler.MiddlewareRequestDeserialization)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}

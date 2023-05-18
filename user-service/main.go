package main

import (
	"BookingPlatform/user-service/handler"
	"BookingPlatform/user-service/repository"
	"BookingPlatform/user-service/service"
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

	fmt.Println("USER")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	userLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)

	userRepository, err := repository.New(timeoutContext, userLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer userRepository.Disconnect(timeoutContext)

	userRepository.Ping()

	userService := service.NewUserService(userRepository, logger)
	userHandler := handler.NewUserHandler(userService, logger)

	userHandler.DatabaseName(timeoutContext)

	router := mux.NewRouter()
	router.Use(userHandler.MiddlewareContentTypeSet)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/insert", userHandler.Insert)
	postRouter.Use(userHandler.MiddlewareUserDeserialization)

	postReservation := router.Methods(http.MethodPost).Subrouter()
	postReservation.HandleFunc("/insertReservation", userHandler.InsertReservation)
	postReservation.Use(userHandler.MiddlewareReservationDeserialization)

	postRequest := router.Methods(http.MethodPost).Subrouter()
	postRequest.HandleFunc("/request/insertReservationRequest", userHandler.InsertReservationRequest)
	postRequest.Use(userHandler.MiddlewareRequestDeserialization)

	postAcceptRequest := router.Methods(http.MethodPost).Subrouter()
	postAcceptRequest.HandleFunc("/request/accept/reservationRequest/{id}", userHandler.AcceptRequest)
	//postAcceptRequest.Use(userHandler.MiddlewareRequestDeserialization)

	logInRouter := router.Methods(http.MethodGet).Subrouter()
	logInRouter.HandleFunc("/login/{username}/{password}", userHandler.GetUserByUsernameAndPassword)

	reservationsRouter := router.Methods(http.MethodGet).Subrouter()
	reservationsRouter.HandleFunc("/getReservations/{id}", userHandler.GetAllReservationsByUser)

	deleteReservaitonRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteReservaitonRouter.HandleFunc("/deleteReservation/{id}", userHandler.DeleteReservation)

	deleteRequestRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRequestRouter.HandleFunc("/delete/request/{id}", userHandler.DeleteRequest)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/update/{id}", userHandler.Update)
	putRouter.Use(userHandler.MiddlewareUserDeserialization)

	apartmentsRouter := router.Methods(http.MethodGet).Subrouter()
	apartmentsRouter.HandleFunc("/apartments/getApartments/{id}", userHandler.FindAllApartmentsByHostId)

	reservationRequestsRouter := router.Methods(http.MethodGet).Subrouter()
	reservationRequestsRouter.HandleFunc("/reservationRequests/requests/getRequests/{id}", userHandler.FindAllReservationRequestsByApartments)

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

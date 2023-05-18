package main

import (
	"BookingPlatform/apartment-service/handler"
	"BookingPlatform/apartment-service/repository"
	"BookingPlatform/apartment-service/service"
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

	fmt.Println("APARTMENT")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	apartmentLogger := log.New(os.Stdout, "[apartment-store] ", log.LstdFlags)

	apartmentRepository, err := repository.New(timeoutContext, apartmentLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer apartmentRepository.Disconnect(timeoutContext)

	apartmentRepository.Ping()

	apartmentService := service.NewApartmentService(apartmentRepository, logger)
	apartmentHandler := handler.NewApartmentHandler(apartmentService, logger)

	apartmentHandler.DatabaseName(timeoutContext)

	router := mux.NewRouter()
	router.Use(apartmentHandler.MiddlewareContentTypeSet)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", apartmentHandler.GetAll)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/insert/{userRole}/{userId}", apartmentHandler.Insert)
	postRouter.Use(apartmentHandler.MiddlewareApartmentDeserialization)

	postPricelistItemRouter := router.Methods(http.MethodPost).Subrouter()
	postPricelistItemRouter.HandleFunc("/insertItem/{apartmentId}/{userRole}", apartmentHandler.InsertPricelistItem)
	postPricelistItemRouter.Use(apartmentHandler.MiddlewarePricelistItemDeserialization)

	filterRouter := router.Methods(http.MethodGet).Subrouter()
	filterRouter.HandleFunc("/filter/{location}/{num}/{start}/{end}", apartmentHandler.FilterApartments)

	getByHostRouter := router.Methods(http.MethodGet).Subrouter()
	getByHostRouter.HandleFunc("/getByHostId/{id}", apartmentHandler.GetAllByHost)

	getOneApartmentRouter := router.Methods(http.MethodGet).Subrouter()
	getOneApartmentRouter.HandleFunc("/getOne/{id}", apartmentHandler.GetOne)

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

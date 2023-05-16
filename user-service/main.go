package main

import (
	"BookingPlatform/user-service/handler"
	"BookingPlatform/user-service/pb"
	"BookingPlatform/user-service/repository"
	"BookingPlatform/user-service/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
)

func main() {
	//var opts []grpc.DialOption
	conn, err := grpc.Dial("localhost:8001", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	}
	defer conn.Close()

	client := pb.NewUserReservationServiceClient(conn)

	//rez, err := client.SayHi(context.Background(), &pb.HiRequest{Message: "VELJA"})

	rez, err := client.GetIfNoReservations(context.Background(), &pb.GetReservationsForUserRequest{GuestId: "646260fa248d4f0da8c81e6a"})
	fmt.Println("USOOO2 " + strconv.FormatBool(rez.Message))

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

	logInRouter := router.Methods(http.MethodGet).Subrouter()
	logInRouter.HandleFunc("/{username}/{password}", userHandler.GetUserByUsernameAndPassword)

	putRouter := router.Methods(http.MethodPatch).Subrouter()
	putRouter.HandleFunc("/update/{id}", userHandler.Update)
	putRouter.Use(userHandler.MiddlewareUserDeserialization)

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

package main

import (
	"BookingPlatform/reservation-service/handler"
	"BookingPlatform/reservation-service/pb"
	"BookingPlatform/reservation-service/repository"
	"BookingPlatform/reservation-service/service"
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
)

type UserReservationService struct {
	pb.UnimplementedUserReservationServiceServer
}

func (s *UserReservationService) SayHi(ctx context.Context, req *pb.HiRequest) (*pb.HiResponse, error) {
	// Logic to fetch user from database or any other service
	//message := "USO1" // Example response
	message := req.Message // Example response

	return &pb.HiResponse{
		Message: message,
	}, nil
}

//func (UnimplementedUserReservationServiceServer) SatHi(context.Context, *HiRequest) (*HiResponse, error) {
//	return nil, status.Errorf(codes.Unimplement

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserReservationServiceServer(grpcServer, &UserReservationService{})
	grpcServer.Serve(lis)

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

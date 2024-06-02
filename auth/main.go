package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"auth.com/handler"
	"auth.com/model"
	"auth.com/proto/auth"
	"auth.com/repo"
	"auth.com/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=soa_auth host=auth-database port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	err = database.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}
	return database
}
func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{Repo: userRepo}
	authHandler := &handler.AuthenticationHandler{AuthService: userService}

	lis, err := net.Listen("tcp", ":90")
	fmt.Println("Running gRPC on port 90")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(lis)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	fmt.Println("Registered gRPC server")

	auth.RegisterAuthServiceServer(grpcServer, authHandler)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
	fmt.Println("Serving gRPC")

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)
	<-stopCh
	grpcServer.Stop()

	// startServer(authHandler)
}

// func startServer(handler *handler.AuthHandler) {
// 	router := mux.NewRouter().StrictSlash(true)

// 	router.HandleFunc("/login", handler.Login).Methods("POST")
// 	router.HandleFunc("/validateToken", handler.ValidateToken).Methods("GET")

// 	println("Server starting")
// 	log.Fatal(http.ListenAndServe(":90", router))
// }

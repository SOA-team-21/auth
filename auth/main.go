package main

import (
	"log"
	"net/http"

	"auth.com/handler"
	"auth.com/model"
	"auth.com/repo"
	"auth.com/service"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=soa_auth host=soa_auth port=5432 sslmode=disable"
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

func startServer(handler *handler.AuthHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/validateToken", handler.ValidateToken).Methods("GET")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":90", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{Repo: userRepo}
	authHandler := &handler.AuthHandler{UserService: userService}

	startServer(authHandler)
}

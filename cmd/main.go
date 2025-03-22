package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gaelzamora/spent-one/config"
	"github.com/gaelzamora/spent-one/internal/adapters/database"
	"github.com/gaelzamora/spent-one/internal/adapters/handlers"
	"github.com/gaelzamora/spent-one/internal/application"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.ConnectDB()

	userRepo := database.NewUserRepository(db)
	authService := application.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	spentRepo := database.NewSpentRepository(db)
	spentService := application(spentRepo)
	spentHandler := handlers.(spentService)

	router := mux.NewRouter()
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(handlers.AuthMiddleware)
	protected.HandleFunc("/me", authHandler.Me).Methods("GET")

	protected.HandleFunc("/spents", authHandler.CreateSpent).Methods("POST")

	fmt.Println("🔵 Servidor corriendo en http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

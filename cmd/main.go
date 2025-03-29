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

	spentRepo := database.NewSpentRepositoryImpl(db)
	spentService := application.NewSpentService(spentRepo)
	spentHandler := handlers.NewSpentHandler(spentService)

	router := mux.NewRouter()
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(handlers.AuthMiddleware)
	protected.HandleFunc("/me", authHandler.Me).Methods("GET")

	// Spents endpoints
	protected.HandleFunc("/spents", spentHandler.CreateSpent).Methods("POST")
	protected.HandleFunc("/spents", spentHandler.GetSpents).Methods("GET")
	protected.HandleFunc("/spents/{id}", spentHandler.GetSpent).Methods("GET")	
	protected.HandleFunc("/spents/{id}", spentHandler.DeleteSpent).Methods("DELETE")
	protected.HandleFunc("/spents/{id}", spentHandler.UpdateSpent).Methods("PATCH")

	fmt.Println("🔵 Servidor corriendo en http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

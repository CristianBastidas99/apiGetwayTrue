package main

import (
	"log"
	"net/http"

	"github.com/CristianBastidas99/apiGetwayGoTrue/handlers"
	"github.com/CristianBastidas99/apiGetwayGoTrue/logger"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Handlers para autenticación y registro
	r.HandleFunc("/auth/login", handlers.HandleLogin).Methods("POST")
	r.HandleFunc("/auth/user", handlers.HandleRegisteUser).Methods("POST")
	r.HandleFunc("/auth/admin", handlers.HandleRegisteAdmin).Methods("POST")
	r.HandleFunc("/api/user/{email}", handlers.HandleGetUser).Methods("GET")
	r.HandleFunc("/api/user", handlers.HandleListUsers).Methods("GET")
	r.HandleFunc("/api/user/password", handlers.HandleChangePassword).Methods("PATCH")
	r.HandleFunc("/api-docs", handlers.HandleAPIDocs).Methods("GET")

	// Handlers para profile
	r.HandleFunc("/api/profile/create", handlers.CreateProfileHandler).Methods("POST")
	r.HandleFunc("/api/profile/{email}", handlers.GetProfileHandler).Methods("GET")
	r.HandleFunc("/api/profile/update/{email}", handlers.UpdateProfileHandler).Methods("POST")
	r.HandleFunc("/api/profile/delete/{email}", handlers.DeleteProfileHandler).Methods("DELETE")
	r.HandleFunc("/api/profile/registration-webhook", handlers.RegistrationWebhookHandler).Methods("POST")

	// Manejadores de estado
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/health/ready", handlers.ReadyCheckHandler).Methods("GET")
	r.HandleFunc("/health/live", handlers.LiveCheckHandler).Methods("GET")

	// Configuración del logger
	logFile, err := logger.SetupLogFile("api_gateway_logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger.SetOutput(logFile)
	r.Use(logger.LogRequests)

	// Configuración del servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("API Gateway running on port 8080")
	log.Fatal(server.ListenAndServe())
}

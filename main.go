package main

import (
	"log"
	"net/http"

	"github.com/CristianBastidas99/apiGetwayGoTrue/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Handlers para autenticación y registro
	r.HandleFunc("/auth/login", handlers.HandleLogin).Methods("POST")
	r.HandleFunc("/auth/user", handlers.HandleRegisteUser).Methods("POST")
	r.HandleFunc("/auth/admin", handlers.HandleRegisteAdmin).Methods("POST")

	// Configuración del logger
	/*logFile, err := logger.SetupLogFile("api_gateway_logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger.SetOutput(logFile)
	r.Use(logger.LogRequests)*/

	// Configuración del servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("API Gateway running on port 8080")
	log.Fatal(server.ListenAndServe())
}
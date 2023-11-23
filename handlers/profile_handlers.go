package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const ProfileAPIURL = "http://profile-service:8080" // Cambiar por la URL real del servicio de seguridad en Java

// CreateProfileHandler maneja la creación de un perfil
func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/create-profile", ProfileAPIURL)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Obtener los datos del cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realizar la solicitud al servicio de perfiles en Java
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio de perfiles al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("CreateProfile request logged")
}

// GetProfileHandler maneja la obtención de un perfil por userID
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	userID := GetEmailHash(email)
	url := fmt.Sprintf("%s/get-profile?userID=%d", ProfileAPIURL, userID)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Realizar la solicitud al servicio de perfiles en Java
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio de perfiles al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("GetProfile request logged")
}

// UpdateProfileHandler maneja la actualización de un perfil por userID
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	userID := GetEmailHash(email)
	url := fmt.Sprintf("%s/update-profile?userID=%d", ProfileAPIURL, userID)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Obtener los datos del cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realizar la solicitud al servicio de perfiles en Java
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio de perfiles al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("UpdateProfile request logged")
}

// DeleteProfileHandler maneja la eliminación de un perfil por userID
func DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	userID := GetEmailHash(email)
	url := fmt.Sprintf("%s/delete-profile?userID=%d", ProfileAPIURL, userID)

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Realizar la solicitud al servicio de perfiles en Java
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio de perfiles al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("DeleteProfile request logged")
}

// RegistrationWebhookHandler maneja el webhook de registro de perfiles
func RegistrationWebhookHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/registration-webhook", ProfileAPIURL)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Obtener los datos del cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realizar la solicitud al servicio de perfiles en Java
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio de perfiles al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("RegistrationWebhook request logged")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/health", ProfileAPIURL) // Reemplaza YourServiceURL por la URL del servicio que deseas verificar

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Realizar la solicitud al servicio para verificar su estado
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("HealthCheck request logged")
}

func ReadyCheckHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/health/ready", ProfileAPIURL) // Reemplaza YourServiceURL por la URL del servicio que deseas verificar

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Realizar la solicitud al servicio para verificar su estado
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("ReadyCheck request logged")
}

func LiveCheckHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/health/live", ProfileAPIURL) // Reemplaza YourServiceURL por la URL del servicio que deseas verificar

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Realizar la solicitud al servicio para verificar su estado
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar la respuesta del servicio al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("LiveCheck request logged")
}

func GetEmailHash(email string) uint32 {
	hasher := sha256.New()
	hasher.Write([]byte(email))
	hashBytes := hasher.Sum(nil)

	// Tomar los primeros 4 byt>es del hash y convertirlos a un entero sin signo de 32 bits
	hashAsInt := binary.BigEndian.Uint32(hashBytes[:4])
	return hashAsInt
}

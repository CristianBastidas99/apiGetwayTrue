package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const securityAPIURL = "http://taller-ms:8080" // Cambiar por la URL real del servicio de seguridad en Java

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/auth/login", securityAPIURL)

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

	// Realizar la solicitud al API de seguridad en Java
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

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("Login request logged")
}

func HandleRegisteUser(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/user/", securityAPIURL)

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

	// Realizar la solicitud al API de seguridad en Java para el registro
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

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("Register user request logged")
}

func HandleRegisteAdmin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/user/admin", securityAPIURL)

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

	// Realizar la solicitud al API de seguridad en Java para el registro
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

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("Register admin request logged")
}

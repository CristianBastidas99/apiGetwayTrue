package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const SecurityAPIURL = "http://taller-ms:8080" // Cambiar por la URL real del servicio de seguridad en Java

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/auth/login", SecurityAPIURL)

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
	url := fmt.Sprintf("%s/api/user/", SecurityAPIURL)

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

	// Crear un mapa para almacenar los datos del JSON de la solicitud
	var requestData map[string]interface{}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener el campo "email" del mapa
	email, ok := requestData["email"].(string)
	if !ok {
		http.Error(w, "Campo 'email' no encontrado en la solicitud JSON o no es una cadena", http.StatusBadRequest)
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

	if resp.StatusCode == http.StatusOK {
		// El usuario se creó correctamente
		// Puedes agregar algún mensaje de éxito o tomar acciones adicionales si es necesario
		userID := GetEmailHash(email)
		url := fmt.Sprintf("%s/registration-webhook", ProfileAPIURL)
		// Enviar al webhook de registro de perfil por correo electrónico
		err = SendToRegistrationWebhook(int(userID), url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Usuario creado correctamente")
		log.Printf("Email del usuario: %s\n", email)
		log.Printf("Email del usuario: %d\n", userID)
	}

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("Register user request logged")
}

func SendToRegistrationWebhook(userID int, webhookURL string) error {
	// Crear el cuerpo de la solicitud para el webhook
	webhookBody := map[string]int{
		"userID": userID,
	}

	// Convertir el cuerpo de la solicitud a JSON
	webhookJSON, err := json.Marshal(webhookBody)
	if err != nil {
		return err
	}

	// Realizar la solicitud al webhook de registro de perfil por correo electrónico
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(webhookJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Manejar la respuesta del webhook si es necesario

	return nil
}

func HandleRegisteAdmin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/user/admin", SecurityAPIURL)

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

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	url := fmt.Sprintf("%s/api/user/%s", SecurityAPIURL, email)

	// Realizar la solicitud GET al API de seguridad en Java para obtener el usuario
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

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("GetUser request logged")
}

func HandleListUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	url := fmt.Sprintf("%s/api/user/?page=%s&size=%s", SecurityAPIURL, page, size)

	// Realizar la solicitud GET al API de seguridad en Java para listar usuarios
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

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("ListUsers request logged")
}

func HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/user/password", SecurityAPIURL)

	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Obtener los datos del cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realizar la solicitud PATCH al API de seguridad en Java para cambiar la contraseña
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

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

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("ChangePassword request logged")
}

func HandleAPIDocs(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api-docs", SecurityAPIURL)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Realizar la solicitud GET al API de seguridad en Java para obtener la documentación
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

	// Retornar la respuesta del API de seguridad al cliente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	log.Println("APIDocs request logged")
}

package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type LogMessage struct {
	AppGener    string `json:"AppGener"`
	Tipo        string `json:"Tipo"`
	ClaseModelo string `json:"ClaseModelo"`
	FechaHora   string `json:"FechaHora"`
	Resumen     string `json:"Resumen"`
	Descripcion string `json:"Descripcion"`
}

var (
	amqpURI = "amqp://guest:guest@rabbitmq:5672/" // Cambia por la URL de conexión de tu servidor RabbitMQ
	queue   = "cola_1"                            // Nombre de la cola
)

func SetupLogFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func SetOutput(file *os.File) {
	log.SetOutput(file)
}

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		start := time.Now()

		logMessage := LogMessage{
			AppGener:    "apigetway",
			Tipo:        r.Method,
			ClaseModelo: r.URL.Path,
			FechaHora:   time.Now().Format("2006-01-02 15:04:05"),
			Resumen:     fmt.Sprintf("Received request: %s %s", r.Method, r.URL.Path),
		}

		if err := sendLogToRabbitMQ(logMessage); err != nil {
			log.Printf("Error sending log message to RabbitMQ: %s", err.Error())
		}

		// Continuar con el manejo de la solicitud
		rec := NewCapturingResponseWriter(w)
		next.ServeHTTP(rec, r)

		// Registro de la respuesta
		logMessage = LogMessage{
			AppGener:    "apigetway",
			Tipo:        r.Method,
			ClaseModelo: r.URL.Path,
			FechaHora:   time.Now().Format("2006-01-02 15:04:05"),
			Resumen:     fmt.Sprintf("Received request: %s %s", r.Method, r.URL.Path),
			Descripcion: r.Method + " " + r.URL.Path + " - " + http.StatusText(rec.Status()),
		}

		if err := sendLogToRabbitMQ(logMessage); err != nil {
			log.Printf("Error sending log message to RabbitMQ: %s", err.Error())
		}

		// Loguear la duración de la solicitud
		log.Printf("Request duration: %s", time.Since(start))

	})
}

// NewCapturingResponseWriter crea un ResponseWriter capturador para rastrear el código de estado de la respuesta.
func NewCapturingResponseWriter(w http.ResponseWriter) *CapturingResponseWriter {
	return &CapturingResponseWriter{w, http.StatusOK}
}

// CapturingResponseWriter implementa http.ResponseWriter para capturar el código de estado de la respuesta.
type CapturingResponseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader registra el código de estado de la respuesta.
func (w *CapturingResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Status devuelve el código de estado capturado.
func (w *CapturingResponseWriter) Status() int {
	return w.status
}

func sendLogToRabbitMQ(logMsg LogMessage) error {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	jsonBody, err := json.Marshal(logMsg)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",    // Intercambio
		queue, // Cola
		false, // Mandar como persistente
		false, // Mandar como publicación inmediata
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

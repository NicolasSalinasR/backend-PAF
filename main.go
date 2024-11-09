package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NicolasSalinasR/Backend-PAF/DB"
	"github.com/gorilla/mux"
)

// Handler de ejemplo para una ruta
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Bienvenido a la API"})
}

func main() {
	// Conectar a la base de datos
	DB.DBconnection()

	// Crear el router
	router := mux.NewRouter()

	// Definir rutas
	router.HandleFunc("/", homeHandler).Methods("GET")

	// Iniciar el servidor
	log.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

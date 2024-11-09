package main

import (
	"log"
	"net/http"

	"github.com/NicolasSalinasR/backend-PAF/DB"
	"github.com/NicolasSalinasR/backend-PAF/controller"
	"github.com/NicolasSalinasR/backend-PAF/service"
	"github.com/gorilla/mux"
)

func main() {
	// Conectar a la base de datos
	DB.DBconnection()
	// Creamos una instancia del servicio
	pafService := service.NewPAFService()

	// Creamos una instancia del controlador
	pafController := controller.NewPAFController(pafService)

	// Configuramos el enrutador
	r := mux.NewRouter()

	// Definir las rutas del controlador
	r.HandleFunc("/paf", pafController.CrearPAF).Methods("POST")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ObtenerPAF).Methods("GET")
	r.HandleFunc("/pafs", pafController.ObtenerTodosPAFs).Methods("GET")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ActualizarPAF).Methods("PUT")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.EliminarPAF).Methods("DELETE")

	// Iniciar el servidor
	log.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

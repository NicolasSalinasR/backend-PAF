package main

import (
	"log"
	"net/http"

	"github.com/NicolasSalinasR/backend-PAF/DB"
	"github.com/NicolasSalinasR/backend-PAF/controller"
	"github.com/gorilla/mux"
)

func main() {
	// Conectar a la base de datos
	DB.DBconnection()

	// Crear una instancia del servicio

	// pafService := service.NewPAFService()

	// Crear una instancia del controlador
	pafController := controller.NewPAFController()

	// Configurar el enrutador
	r := mux.NewRouter()

	// Definir las rutas para los PAFs
	r.HandleFunc("/paf", pafController.CrearPAF).Methods("POST")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ObtenerPAF).Methods("GET")
	r.HandleFunc("/pafs", pafController.ObtenerTodosPAFs).Methods("GET")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ActualizarPAF).Methods("PUT")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.EliminarPAF).Methods("DELETE")
	r.HandleFunc("/pafs/profesor/{nombreProfesor}", pafController.ObtenerPAFsPorNombreProfesor).Methods("GET")

	// Iniciar el servidor
	log.Println("Servidor escuchando en el puerto 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}

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
	profesorService := service.NewProfesorService(DB.DB)

	// Creamos una instancia del controlador
	pafController := controller.NewPAFController(pafService)
	profesorController := controller.NewProfesorController(profesorService)

	// Configuramos el enrutador
	r := mux.NewRouter()

	// Definir las rutas del controlador para PAF
	r.HandleFunc("/paf", pafController.CrearPAF).Methods("POST")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ObtenerPAF).Methods("GET")
	r.HandleFunc("/pafs", pafController.ObtenerTodosPAFs).Methods("GET")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ActualizarPAF).Methods("PUT")
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.EliminarPAF).Methods("DELETE")

	// Definir las rutas del controlador para Profesores
	r.HandleFunc("/profesores", profesorController.CrearProfesor).Methods("POST")
	r.HandleFunc("/profesores/{id:[0-9]+}", profesorController.ObtenerProfesor).Methods("GET")
	r.HandleFunc("/profesores", profesorController.ObtenerTodosProfesores).Methods("GET")
	r.HandleFunc("/profesores/{id:[0-9]+}", profesorController.ActualizarProfesor).Methods("PUT")
	r.HandleFunc("/profesores/{id:[0-9]+}", profesorController.EliminarProfesor).Methods("DELETE")

	// Iniciar el servidor
	log.Println("Servidor escuchando en el puerto 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}

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
	//Filtro Por Id PAF
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ObtenerPAF).Methods("GET")
	//Obtiene todas las PAFS
	r.HandleFunc("/pafs", pafController.ObtenerTodosPAFs).Methods("GET")
	//Actualiza la PAF del id ingresado
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.ActualizarPAF).Methods("PUT")
	//Elimina una PAF
	r.HandleFunc("/paf/{id:[0-9]+}", pafController.EliminarPAF).Methods("DELETE")

	// Ruta para filtrar por nombre del profesor
	r.HandleFunc("/pafs/buscarNombre", pafController.ObtenerPAFsPorNombreProfesor).Methods("GET")

	// Iniciar el servidor
	log.Println("Servidor escuchando en el puerto 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}

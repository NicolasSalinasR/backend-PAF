package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NicolasSalinasR/backend-PAF/models"
	"github.com/NicolasSalinasR/backend-PAF/service"
	"github.com/gorilla/mux"
)

// PAFController define las operaciones HTTP sobre PAF
type PAFController struct {
	PafService *service.PAFService
}

// NewPAFController crea una nueva instancia del controlador
func NewPAFController(pafService *service.PAFService) *PAFController {
	return &PAFController{PafService: pafService}
}

// CrearPAF maneja la creación de un nuevo PAF
func (c *PAFController) CrearPAF(w http.ResponseWriter, r *http.Request) {
	var paf models.PAF
	// Decodificar el cuerpo de la solicitud (JSON)
	if err := json.NewDecoder(r.Body).Decode(&paf); err != nil {
		http.Error(w, "Error al leer los datos del cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para crear el PAF
	createdPAF, err := c.PafService.CrearPAF(&paf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar una respuesta con el PAF creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdPAF); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// ObtenerPAF maneja la obtención de un PAF por ID
func (c *PAFController) ObtenerPAF(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro ID de la URL
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para obtener el PAF
	paf, err := c.PafService.ObtenerPAF(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Enviar la respuesta con el PAF
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(paf); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// ObtenerTodosPAFs maneja la obtención de todos los PAFs
func (c *PAFController) ObtenerTodosPAFs(w http.ResponseWriter, r *http.Request) {
	// Llamar al servicio para obtener todos los PAFs
	pafs, err := c.PafService.ObtenerTodosPAFs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar la respuesta con la lista de PAFs
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pafs); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// ActualizarPAF maneja la actualización de un PAF por ID
func (c *PAFController) ActualizarPAF(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var paf models.PAF
	// Decodificar el cuerpo de la solicitud (JSON)
	if err := json.NewDecoder(r.Body).Decode(&paf); err != nil {
		http.Error(w, "Error al leer los datos del cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para actualizar el PAF
	updatedPAF, err := c.PafService.ActualizarPAF(uint(id), &paf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar la respuesta con el PAF actualizado
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedPAF); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// EliminarPAF maneja la eliminación de un PAF por ID
func (c *PAFController) EliminarPAF(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para eliminar el PAF
	if err := c.PafService.EliminarPAF(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar una respuesta vacía indicando que el PAF fue eliminado
	w.WriteHeader(http.StatusNoContent)
}

// ObtenerPAFsPorNombreProfesor maneja la solicitud GET para obtener PAFs filtrados por nombre del profesor
func (c *PAFController) ObtenerPAFsPorNombreProfesor(w http.ResponseWriter, r *http.Request) {
	// Obtener el nombre del profesor del query string
	nombreProfesor := r.URL.Query().Get("nombre_profesor")
	if nombreProfesor == "" {
		http.Error(w, "El parámetro 'nombre_profesor' es obligatorio", http.StatusBadRequest)
		return
	}

	// Obtener los PAFs filtrados
	pafs, err := c.PafService.ObtenerPAFsPorNombreProfesor(nombreProfesor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respondemos con los resultados en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pafs)
}

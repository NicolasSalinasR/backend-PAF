package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NicolasSalinasR/backend-PAF/models"
	"github.com/NicolasSalinasR/backend-PAF/service"
	"github.com/gorilla/mux"
)

// PAFController maneja las solicitudes HTTP para los PAFs
type PAFController struct {
	PAFService *service.PAFService
}

// NewPAFController crea una nueva instancia de PAFController
func NewPAFController() *PAFController {
	return &PAFController{
		PAFService: service.NewPAFService(),
	}
}

// CrearPAF maneja la creación de un nuevo PAF
func (c *PAFController) CrearPAF(w http.ResponseWriter, r *http.Request) {
	var paf models.PAF
	// Decodificar el cuerpo de la solicitud en un objeto PAF
	if err := json.NewDecoder(r.Body).Decode(&paf); err != nil {
		http.Error(w, "Error al leer los datos del PAF", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para crear el PAF
	createdPAF, err := c.PAFService.CrearPAF(&paf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	// Responder con el PAF creado
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPAF)
}

// ObtenerPAF maneja la obtención de un PAF por ID
func (c *PAFController) ObtenerPAF(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del PAF desde los parámetros de la URL
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para obtener el PAF
	paf, err := c.PAFService.ObtenerPAF(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Responder con el PAF encontrado
	json.NewEncoder(w).Encode(paf)
}

// ObtenerTodosPAFs maneja la obtención de todos los PAFs
func (c *PAFController) ObtenerTodosPAFs(w http.ResponseWriter, r *http.Request) {
	// Llamar al servicio para obtener todos los PAFs
	pafs, err := c.PAFService.ObtenerTodosPAFs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con la lista de PAFs
	json.NewEncoder(w).Encode(pafs)
}

// ActualizarPAF maneja la actualización de un PAF
func (c *PAFController) ActualizarPAF(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del PAF desde los parámetros de la URL
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var paf models.PAF
	// Decodificar el cuerpo de la solicitud en un objeto PAF
	if err := json.NewDecoder(r.Body).Decode(&paf); err != nil {
		http.Error(w, "Error al leer los datos del PAF", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para actualizar el PAF
	updatedPAF, err := c.PAFService.ActualizarPAF(uint(id), &paf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Responder con el PAF actualizado
	json.NewEncoder(w).Encode(updatedPAF)
}

// EliminarPAF maneja la eliminación de un PAF por ID
func (c *PAFController) EliminarPAF(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del PAF desde los parámetros de la URL
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para eliminar el PAF
	err = c.PAFService.EliminarPAF(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusNoContent)
}

// ObtenerPAFsPorNombreProfesor maneja la obtención de los PAFs por el nombre del profesor
func (c *PAFController) ObtenerPAFsPorNombreProfesor(w http.ResponseWriter, r *http.Request) {
	nombreProfesor := mux.Vars(r)["nombreProfesor"]

	// Llamar al servicio para obtener los PAFs por el nombre del profesor
	pafs, err := c.PAFService.ObtenerPAFsPorNombreProfesor(nombreProfesor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con los PAFs encontrados
	json.NewEncoder(w).Encode(pafs)
}

package controller

import (
	"fmt"
	"net/http"

	"github.com/NicolasSalinasR/backend-PAF/models"
	"github.com/NicolasSalinasR/backend-PAF/service"
	"github.com/gin-gonic/gin"
)

// ProfesorController define las operaciones del controlador de los profesores
type ProfesorController struct {
	ProfesorService *service.ProfesorService
}

// NewProfesorController crea una nueva instancia del ProfesorController
func NewProfesorController(profesorService *service.ProfesorService) *ProfesorController {
	return &ProfesorController{ProfesorService: profesorService}
}

// CrearProfesor maneja la creación de un nuevo Profesor
func (ctrl *ProfesorController) CrearProfesor(c *gin.Context) {
	var profesor models.Profesor

	// Bind JSON a la estructura Profesor
	if err := c.ShouldBindJSON(&profesor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para crear el Profesor
	createdProfesor, err := ctrl.ProfesorService.CrearProfesor(&profesor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al crear el profesor: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, createdProfesor)
}

// ObtenerProfesor maneja la obtención de un Profesor por su ID
func (ctrl *ProfesorController) ObtenerProfesor(c *gin.Context) {
	id := c.Param("id") // Obtener el parámetro ID de la URL
	var profesor models.Profesor

	// Convertir el ID a un tipo adecuado
	if err := c.ShouldBindUri(&profesor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al servicio para obtener el Profesor por ID
	queriedProfesor, err := ctrl.ProfesorService.ObtenerProfesor(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Profesor con ID %s no encontrado", id)})
		return
	}

	c.JSON(http.StatusOK, queriedProfesor)
}

// ObtenerTodosProfesores maneja la obtención de todos los Profesores
func (ctrl *ProfesorController) ObtenerTodosProfesores(c *gin.Context) {
	profesores, err := ctrl.ProfesorService.ObtenerTodosProfesores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los profesores"})
		return
	}

	c.JSON(http.StatusOK, profesores)
}

// ActualizarProfesor maneja la actualización de un Profesor existente
func (ctrl *ProfesorController) ActualizarProfesor(c *gin.Context) {
	id := c.Param("id") // Obtener el parámetro ID de la URL
	var profesor models.Profesor

	// Bind JSON a la estructura Profesor
	if err := c.ShouldBindJSON(&profesor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamar al servicio para actualizar el Profesor
	updatedProfesor, err := ctrl.ProfesorService.ActualizarProfesor(id, &profesor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al actualizar el profesor: %v", err)})
		return
	}

	c.JSON(http.StatusOK, updatedProfesor)
}

// EliminarProfesor maneja la eliminación de un Profesor por su ID
func (ctrl *ProfesorController) EliminarProfesor(c *gin.Context) {
	id := c.Param("id") // Obtener el parámetro ID de la URL

	// Llamar al servicio para eliminar el Profesor
	if err := ctrl.ProfesorService.EliminarProfesor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al eliminar el profesor: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profesor eliminado exitosamente"})
}

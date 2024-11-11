package service

import (
	"fmt"

	"github.com/NicolasSalinasR/backend-PAF/models"
	"gorm.io/gorm"
)

// ProfesorService define las operaciones de la lógica de negocio sobre Profesor
type ProfesorService struct {
	DB *gorm.DB
}

// NewProfesorService crea una nueva instancia de ProfesorService
func NewProfesorService(db *gorm.DB) *ProfesorService {
	return &ProfesorService{DB: db}
}

// CrearProfesor crea un nuevo registro de Profesor en la base de datos
func (s *ProfesorService) CrearProfesor(profesor *models.Profesor) (*models.Profesor, error) {
	if err := s.DB.Create(&profesor).Error; err != nil {
		return nil, err
	}
	return profesor, nil
}

// ObtenerProfesor obtiene un Profesor por ID
func (s *ProfesorService) ObtenerProfesor(id uint) (*models.Profesor, error) {
	var profesor models.Profesor
	if err := s.DB.First(&profesor, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Profesor con ID %d no encontrado", id)
		}
		return nil, err
	}
	return &profesor, nil
}

// ObtenerTodosProfesores obtiene todos los registros de Profesor
func (s *ProfesorService) ObtenerTodosProfesores() ([]models.Profesor, error) {
	var profesores []models.Profesor
	if err := s.DB.Find(&profesores).Error; err != nil {
		return nil, err
	}
	return profesores, nil
}

// ActualizarProfesor actualiza los detalles de un Profesor
func (s *ProfesorService) ActualizarProfesor(id uint, profesor *models.Profesor) (*models.Profesor, error) {
	var existingProfesor models.Profesor
	if err := s.DB.First(&existingProfesor, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Profesor con ID %d no encontrado", id)
		}
		return nil, err
	}

	existingProfesor.NombreProfesor = profesor.NombreProfesor
	existingProfesor.RutProfesor = profesor.RutProfesor
	existingProfesor.Contraseña = profesor.Contraseña
	existingProfesor.Correo = profesor.Correo

	if err := s.DB.Save(&existingProfesor).Error; err != nil {
		return nil, err
	}

	return &existingProfesor, nil
}

// EliminarProfesor elimina un registro de Profesor por ID
func (s *ProfesorService) EliminarProfesor(id uint) error {
	if err := s.DB.Delete(&models.Profesor{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ObtenerProfesorPorNombre obtiene Profesores por nombre
func (s *ProfesorService) ObtenerProfesorPorNombre(nombre string) ([]models.Profesor, error) {
	var profesores []models.Profesor
	if err := s.DB.Where("nombre_profesor = ?", nombre).Find(&profesores).Error; err != nil {
		return nil, err
	}
	return profesores, nil
}

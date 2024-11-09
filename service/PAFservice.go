package service

import (
	"fmt"

	"github.com/NicolasSalinasR/backend-PAF/DB"
	"github.com/NicolasSalinasR/backend-PAF/models"
	"gorm.io/gorm"
)

// PAFService define las operaciones de la lógica de negocio sobre PAF
type PAFService struct{}

// NewPAFService crea una nueva instancia de PAFService
func NewPAFService() *PAFService {
	return &PAFService{}
}

// CrearPAF crea un nuevo registro de PAF en la base de datos
func (s *PAFService) CrearPAF(paf *models.PAF) (*models.PAF, error) {
	// Validar que no exista un PAF con el mismo código de asignatura (ejemplo de validación)
	var existingPAF models.PAF
	if err := DB.DB.Where("codigo_asignatura = ?", paf.CodigoAsignatura).First(&existingPAF).Error; err == nil {
		return nil, fmt.Errorf("ya existe un PAF con el código de asignatura %s", paf.CodigoAsignatura)
	}

	// Crear el PAF en la base de datos
	if err := DB.DB.Create(&paf).Error; err != nil {
		return nil, err
	}

	return paf, nil
}

// ObtenerPAF por ID
func (s *PAFService) ObtenerPAF(id uint) (*models.PAF, error) {
	var paf models.PAF
	if err := DB.DB.First(&paf, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("PAF con ID %d no encontrado", id)
		}
		return nil, err
	}
	return &paf, nil
}

// ObtenerTodosPAFs obtiene todos los registros de PAF
func (s *PAFService) ObtenerTodosPAFs() ([]models.PAF, error) {
	var pAFs []models.PAF
	if err := DB.DB.Find(&pAFs).Error; err != nil {
		return nil, err
	}
	return pAFs, nil
}

// ActualizarPAF actualiza los detalles de un PAF
func (s *PAFService) ActualizarPAF(id uint, paf *models.PAF) (*models.PAF, error) {
	var existingPAF models.PAF
	if err := DB.DB.First(&existingPAF, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("PAF con ID %d no encontrado", id)
		}
		return nil, err
	}

	// Actualizamos los campos del PAF
	existingPAF.NombreProfesor = paf.NombreProfesor
	existingPAF.Cupo = paf.Cupo
	existingPAF.Grupo = paf.Grupo
	existingPAF.Fecha = paf.Fecha
	existingPAF.Etapa = paf.Etapa
	existingPAF.CodigoAsignatura = paf.CodigoAsignatura

	// Guardamos los cambios
	if err := DB.DB.Save(&existingPAF).Error; err != nil {
		return nil, err
	}

	return &existingPAF, nil
}

// EliminarPAF elimina un registro de PAF por ID
func (s *PAFService) EliminarPAF(id uint) error {
	if err := DB.DB.Delete(&models.PAF{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ObtenerPAFsPorNombreProfesor obtiene todos los registros de PAF filtrados por el nombre del profesor
func (s *PAFService) ObtenerPAFsPorNombreProfesor(nombreProfesor string) ([]models.PAF, error) {
	var pAFs []models.PAF
	if err := DB.DB.Where("nombre_profesor = ?", nombreProfesor).Find(&pAFs).Error; err != nil {
		return nil, err
	}
	return pAFs, nil
}

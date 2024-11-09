package service

import (
	"fmt"

	"github.com/NicolasSalinasR/backend-PAF/model"
	"gorm.io/gorm"
)

// PAFService define las operaciones de la lógica de negocio sobre PAF
type PAFService struct {
	DB *gorm.DB
}

// NewPAFService crea una nueva instancia de PAFService
func NewPAFService(db *gorm.DB) *PAFService {
	return &PAFService{DB: db}
}

// CrearPAF crea un nuevo registro de PAF en la base de datos
func (s *PAFService) CrearPAF(paf *model.PAF) (*model.PAF, error) {
	// Validar que no exista un PAF con el mismo código de asignatura (ejemplo de validación)
	var existingPAF model.PAF
	if err := s.DB.Where("codigo_asignatura = ?", paf.CodigoAsignatura).First(&existingPAF).Error; err == nil {
		return nil, fmt.Errorf("ya existe un PAF con el código de asignatura %s", paf.CodigoAsignatura)
	}

	// Crear el PAF en la base de datos
	if err := s.DB.Create(&paf).Error; err != nil {
		return nil, err
	}

	return paf, nil
}

// ObtenerPAF por ID
func (s *PAFService) ObtenerPAF(id uint) (*model.PAF, error) {
	var paf model.PAF
	if err := s.DB.First(&paf, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("PAF con ID %d no encontrado", id)
		}
		return nil, err
	}
	return &paf, nil
}

// ObtenerTodosPAFs obtiene todos los registros de PAF
func (s *PAFService) ObtenerTodosPAFs() ([]model.PAF, error) {
	var pAFs []model.PAF
	if err := s.DB.Find(&pAFs).Error; err != nil {
		return nil, err
	}
	return pAFs, nil
}

// ActualizarPAF actualiza los detalles de un PAF
func (s *PAFService) ActualizarPAF(id uint, paf *model.PAF) (*model.PAF, error) {
	var existingPAF model.PAF
	if err := s.DB.First(&existingPAF, id).Error; err != nil {
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
	if err := s.DB.Save(&existingPAF).Error; err != nil {
		return nil, err
	}

	return &existingPAF, nil
}

// EliminarPAF elimina un registro de PAF por ID
func (s *PAFService) EliminarPAF(id uint) error {
	if err := s.DB.Delete(&model.PAF{}, id).Error; err != nil {
		return err
	}
	return nil
}

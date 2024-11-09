package models

import (
	"time"

	"gorm.io/gorm"
)

type PAF struct {
	gorm.Model                 // Incluye ID, CreatedAt, UpdatedAt, DeletedAt
	NombreProfesor   string    `gorm:"type:text;not null"` // Nombre del profesor usando text
	Cupo             int       `gorm:"not null"`           // Cupo (número entero)
	Grupo            string    `gorm:"type:text;not null"` // Grupo usando text
	Fecha            time.Time `gorm:"not null"`           // Fecha
	Etapa            string    `gorm:"type:text;not null"` // Etapa usando text
	RutProfesor      string    `gorm:"type:text;not null"` // Etapa usando text
	CodigoAsignatura string    `gorm:"type:text;not null"` // Código de asignatura usando text
}

package models

import (
	"gorm.io/gorm"
)

type Profesor struct {
	gorm.Model            // Incluye ID, CreatedAt, UpdatedAt, DeletedAt
	NombreProfesor string `gorm:"type:text;not null"` // Nombre del profesor usando text
	RutProfesor    string `gorm:"type:text;not null"` // Etapa usando text
	Contraseña     string `gorm:"type:text;not null"` // Etapa usando text
	Correo         string `gorm:"type:text;not null"` // Etapa usando text

}

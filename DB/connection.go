package DB

import (
	"fmt"
	"log"

	"github.com/NicolasSalinasR/backend-PAF/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=conan2084 dbname=PAF port=5432 sslmode=disable"
var DB *gorm.DB

// DBconnection establece la conexión a la base de datos y la asigna a la variable global DB.
func DBconnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Migrar el modelo (crear la tabla si no existe)
	err = DB.AutoMigrate(models.PAF{})
	if err != nil {
		panic("failed to migrate the database")
	}
	fmt.Println("Conexión a la base de datos exitosa.")

	// Realizar la migración automática para crear las tablas
	err = DB.AutoMigrate(&models.Profesor{})
	if err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	} else {
		log.Println("Migración realizada correctamente")
	}

}

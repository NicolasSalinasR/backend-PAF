package DB

import (
	"fmt"
	"log"

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
	fmt.Println("Conexión a la base de datos exitosa.")
}

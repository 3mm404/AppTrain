package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var MYDB *gorm.DB

func ConectionDB() {
	// Cargar variables del .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ Error cargando el archivo .env:", err)
		return
	}

	// Obtener la URL de conexión desde el .env
	dsn := os.Getenv("DATABASE_URL")

	// Si no se puede conectar con la URL de conexión pooling, intentar con la URL directa
	if dsn == "" {
		dsn = os.Getenv("DIRECT_URL")
	}

	// Conectar con GORM
	MYDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("❌ Error al conectar a la base de datos:", err)
	} else {
		fmt.Println("✅ Conectado a la base de datos de Supabase con .env")
	}
}

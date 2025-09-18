package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var MYDB *gorm.DB

func Connect() {
	//sslmode=require
	dsn := "host=localhost user=postgres password=emmanuel dbname=dbtranify port=5000 TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}

	MYDB = db
	fmt.Println("✅ Conexión exitosa a Supabase PostgreSQL")
}

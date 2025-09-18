package database

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"os"

	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
)

func Migrate() error {
	tables := []interface{}{
		&models.User{},
		&models.Gym{},
		&models.MembershipType{},
		&models.UserMembership{},
		&models.Payment{},
		&models.Attendance{},
		&models.GymEmployee{},
	}

	// Leer variable DROP_TABLES del entorno
	if os.Getenv("DROP_TABLES") == "true" {
		for _, table := range tables {
			if err := config.MYDB.Migrator().DropTable(table); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("error al eliminar tabla: %v", err)
			}
		}
		fmt.Println("Tablas eliminadas")
	}

	if err := config.MYDB.AutoMigrate(tables...); err != nil {
		return fmt.Errorf("error al crear tablas: %v", err)
	}

	fmt.Println("Migración exitosa")
	return nil
}

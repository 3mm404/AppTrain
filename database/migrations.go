package database

import (
	"TrainApp/config"
	"TrainApp/models"
	"fmt"
)

func Migrate() error {

	error := config.MYDB.Error

	config.MYDB.AutoMigrate(
		&models.User{},
		&models.Gym{},
		&models.MembershipType{},
		&models.UserMembership{},
		&models.Payment{},
		&models.Attendance{},
		&models.GymEmployee{},
	)

	if error != nil {
		fmt.Println("Error al migrar la base de datos")
		return error
	} else {
		fmt.Println("Migracion exitosa")
		return nil
	}

}

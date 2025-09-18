package gym_services

import (
	"fmt"

	"github.com/3mm404/gymgo/config"
	dtos "github.com/3mm404/gymgo/gym/dtos"
	"github.com/3mm404/gymgo/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterGym(gymDTO dtos.RegisterGymDTO) (models.Gym, error) {
	var existing models.Gym

	// Verificar teléfono duplicado
	if err := config.MYDB.Where("telefono = ?", gymDTO.Telefono).First(&existing).Error; err == nil {
		return models.Gym{}, fmt.Errorf("el teléfono ya está registrado")
	} else if err != gorm.ErrRecordNotFound {
		return models.Gym{}, err
	}

	// Verificar correo duplicado
	if err := config.MYDB.Where("email = ?", gymDTO.Email).First(&existing).Error; err == nil {
		return models.Gym{}, fmt.Errorf("el correo electrónico ya está registrado")
	} else if err != gorm.ErrRecordNotFound {
		return models.Gym{}, err
	}

	// Hashear contraseña
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(gymDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.Gym{}, err
	}

	// Crear nuevo gimnasio
	newGym := models.Gym{
		Nombre:    gymDTO.Nombre,
		Direccion: gymDTO.Direccion,
		Telefono:  gymDTO.Telefono,
		Email:     gymDTO.Email,
		Foto:      gymDTO.Foto,
		Latitud:   gymDTO.Latitud,
		Longitud:  gymDTO.Longitud,
		Aprobado:  gymDTO.Aprobado,
		CreadoEn:  gymDTO.CreadoEn,
		Password:  string(hashedPass),
	}

	// Guardar en la base de datos
	if err := config.MYDB.Create(&newGym).Error; err != nil {
		return models.Gym{}, err
	}

	return newGym, nil
}

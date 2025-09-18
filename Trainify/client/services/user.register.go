package userservices

import (
	"fmt"

	dtos_user "github.com/3mm404/gymgo/client/dtos"
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(user dtos_user.RegisterUserDTO) (models.User, error) {
	var existing models.User

	// Verificar duplicados (email o teléfono)
	if err := config.MYDB.
		Where("email = ? OR telefono = ?", user.Email, user.Telefono).
		First(&existing).Error; err == nil {
		return models.User{}, fmt.Errorf("el correo o teléfono ya están registrados")
	} else if err != gorm.ErrRecordNotFound {
		return models.User{}, err
	}

	// Hashear contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	// Crear nuevo usuario
	newUser := models.User{
		Nombre:   user.Nombre,
		Email:    user.Email,
		Telefono: user.Telefono,
		Password: string(hashedPassword),
	}

	if err := config.MYDB.Create(&newUser).Error; err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

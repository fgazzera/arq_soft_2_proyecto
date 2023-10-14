package user

import (
	"usuario_reserva-api/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertUser(user models.User) models.User {
	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Usuario Creado: ", user.ID)
	return user
}

func GetUserById(id int) models.User {
	var user models.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("Usuario: ", user)

	return user
}

func GetUserByUsername(username string) models.User {
	var user models.User

	Db.Where("user_name = ?", username).First(&user)
	log.Debug("Usuario: ", user)

	return user
}

func GetUserByEmail(email string) models.User {
	var user models.User

	Db.Where("email = ?", email).First(&user)
	log.Debug("Usuario: ", user)

	return user
}

func GetUsers() models.Users {
	var users models.Users

	Db.Find(&users)
	log.Debug("Usuarios: ", users)

	return users
}

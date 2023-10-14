package reserva

import (
	"usuario_reserva-api/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertReserva(reserva models.Reserva) models.Reserva {
	result := Db.Create(&reserva)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Reserva Creada: ", reserva.ID)
	return reserva
}

func GetReservaById(id int) models.Reserva {
	var reserva models.Reserva

	// Db.Where("id = ?", id).Preload("Hotel").Preload("User").First(&reserva)
	Db.Where("id = ?", id).Preload("User").First(&reserva)
	log.Debug("Reserva: ", reserva)

	return reserva
}

func GetReservasById(id int) models.Reservas {
	var reservas models.Reservas

	// Db.Where("user_id = ?", id).Preload("Hotel").Preload("User").Find(&reservas)
	Db.Where("user_id = ?", id).Preload("User").Find(&reservas)
	log.Debug("Reservas: ", reservas)

	return reservas
}

func GetReservas() models.Reservas {
	var reservas models.Reservas

	// Db.Preload("Hotel").Preload("User").Find(&reservas)
	Db.Preload("User").Find(&reservas)
	log.Debug("Reservas: ", reservas)

	return reservas
}

func GetDisponibilidad(id int) models.Reservas {
	var reservas models.Reservas

	// Db.Where("hotel_id = ?", id).Preload("Hotel").Preload("User").Find(&reservas)
	Db.Where("hotel_id = ?", id).Preload("User").Find(&reservas)
	log.Debug("Reservas: ", reservas)

	return reservas
}

func GetReservasByDate() models.Reservas {
	var reservas models.Reservas

	// Db.Preload("Hotel").Preload("User").Find(&reservas)
	Db.Preload("User").Find(&reservas)
	log.Debug("Reservas: ", reservas)

	return reservas
}

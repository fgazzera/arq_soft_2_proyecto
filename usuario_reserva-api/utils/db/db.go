package db

import (
	reservaDao "usuario_reserva-api/daos/reserva"
	userDao "usuario_reserva-api/daos/user"

	"usuario_reserva-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "arqsoft2"
	DBUser := "root"
	DBPass := ""
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("La conexión no se pudo abrir")
		log.Fatal(err)
	} else {
		log.Info("Conexión establecida")
	}

	userDao.Db = db
	reservaDao.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Reserva{})

	log.Info("Finalizacion de las tablas de la base de datos de migracion")
}

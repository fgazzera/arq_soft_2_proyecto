package model

type Reserva struct {
	ID int `gorm:"primaryKey"`

	Hotel   Hotel `gorm:"foreignkey:HotelID"`
	HotelID int

	User   User `gorm:"foreignkey:UserID"`
	UserID int

	AnioInicio int `gorm:"type:varchar(10)"`
	AnioFinal  int `gorm:"type:varchar(10)"`
	MesInicio  int `gorm:"type:varchar(10)"`
	MesFinal   int `gorm:"type:varchar(10)"`
	DiaInicio  int `gorm:"type:varchar(10)"`
	DiaFinal   int `gorm:"type:varchar(10)"`
	Dias       int `gorm:"type:varchar(2)"`
}

type Reservas []Reserva

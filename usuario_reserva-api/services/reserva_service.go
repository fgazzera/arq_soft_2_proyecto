package services

import (
	reservaDao "usuario_reserva-api/daos/reserva"

	dtos "usuario_reserva-api/dtos"
	models "usuario_reserva-api/models"
	e "usuario_reserva-api/utils/errors"
)

type reservaService struct{}

type reservaServiceInterface interface {
	InsertReserva(reservaDto dtos.ReservaDto) (dtos.ReservaDto, e.ApiError)
	GetReservaById(id int) (dtos.ReservaDto, e.ApiError)
	GetReservasById(id int) (dtos.ReservasDto, e.ApiError)
	GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int)
	GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dtos.ReservasDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) InsertReserva(reservaDto dtos.ReservaDto) (dtos.ReservaDto, e.ApiError) {

	var reserva models.Reserva
	//var hotel models.Hotel
	var user models.User

	reserva.AnioInicio = reservaDto.AnioInicio
	reserva.AnioFinal = reservaDto.AnioFinal
	reserva.MesInicio = reservaDto.MesInicio
	reserva.MesFinal = reservaDto.MesFinal
	reserva.DiaInicio = reservaDto.DiaInicio
	reserva.DiaFinal = reservaDto.DiaFinal
	reserva.Dias = reservaDto.Dias

	reserva.HotelID = reservaDto.HotelID
	reserva.UserID = reservaDto.UserID

	//reserva.Hotel = hotel
	reserva.User = user

	reserva = reservaDao.InsertReserva(reserva)

	reservaDto.ID = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetReservaById(id int) (dtos.ReservaDto, e.ApiError) {

	var reserva models.Reserva = reservaDao.GetReservaById(id)
	var reservaDto dtos.ReservaDto

	if reserva.ID == 0 {
		return reservaDto, e.NewBadRequestApiError("Reserva No Encontrada")
	}

	reservaDto.ID = reserva.ID
	//reservaDto.HotelID = reserva.Hotel.ID
	reservaDto.UserID = reserva.User.ID
	reservaDto.AnioInicio = reserva.AnioInicio
	reservaDto.AnioFinal = reserva.AnioFinal
	reservaDto.MesInicio = reserva.MesInicio
	reservaDto.MesFinal = reserva.MesFinal
	reservaDto.DiaInicio = reserva.DiaInicio
	reservaDto.DiaFinal = reserva.DiaFinal
	reservaDto.Dias = reserva.Dias

	return reservaDto, nil
}

func (s *reservaService) GetReservasById(id int) (dtos.ReservasDto, e.ApiError) {

	var reservas models.Reservas = reservaDao.GetReservasById(id)
	var reservasDto dtos.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dtos.ReservaDto

		if reserva.UserID == 0 {
			return reservasDto, e.NewBadRequestApiError("Reservas No Encontradas")
		}

		reservaDto.ID = reserva.ID
		//reservaDto.HotelID = reserva.Hotel.ID
		reservaDto.UserID = reserva.User.ID
		reservaDto.AnioInicio = reserva.AnioInicio
		reservaDto.AnioFinal = reserva.AnioFinal
		reservaDto.MesInicio = reserva.MesInicio
		reservaDto.MesFinal = reserva.MesFinal
		reservaDto.DiaInicio = reserva.DiaInicio
		reservaDto.DiaFinal = reserva.DiaFinal
		reservaDto.Dias = reserva.Dias

		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int) {

	// var reservas models.Reservas = reservaDao.GetDisponibilidad(id)
	// cambiar: hay que traer la info del hotel desde la api de ficha de hotel

	// var hotel models.Hotel = hotelDao.GetHotelById(id)

	// El resto usa la info del hotel, la logica esta bien

	/*
		disponibilidad = hotel.Cant_Hab

		for _, reserva := range reservas {
			if reserva.HotelID == id && reserva.AnioFinal >= AnioInicio && reserva.AnioInicio <= AnioFinal && reserva.MesFinal >= MesInicio && reserva.MesInicio <= MesFinal && reserva.DiaFinal >= DiaInicio && reserva.DiaInicio <= DiaFinal {
				disponibilidad --
			}
		}
	*/

	return 0
}

func (s *reservaService) GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dtos.ReservasDto, e.ApiError) {

	var reservas models.Reservas = reservaDao.GetReservasByDate()
	var reservasDto dtos.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dtos.ReservaDto

		if reserva.AnioFinal >= AnioInicio && reserva.AnioInicio <= AnioFinal && reserva.MesFinal >= MesInicio && reserva.MesInicio <= MesFinal && reserva.DiaFinal >= DiaInicio && reserva.DiaInicio <= DiaFinal {
			reservaDto.ID = reserva.ID
			//reservaDto.HotelID = reserva.Hotel.ID
			reservaDto.UserID = reserva.User.ID
			reservaDto.AnioInicio = reserva.AnioInicio
			reservaDto.AnioFinal = reserva.AnioFinal
			reservaDto.MesInicio = reserva.MesInicio
			reservaDto.MesFinal = reserva.MesFinal
			reservaDto.DiaInicio = reserva.DiaInicio
			reservaDto.DiaFinal = reserva.DiaFinal
			reservaDto.Dias = reserva.Dias

			reservasDto = append(reservasDto, reservaDto)
		}
	}

	return reservasDto, nil
}

package services

import (
	hotelClient "backend/clients/hotel"
	reservaClient "backend/clients/reserva"

	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type reservaService struct{}

type reservaServiceInterface interface {
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasById(id int) (dto.ReservasDto, e.ApiError)
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int)
	GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva
	var hotel model.Hotel
	var cliente model.Cliente

	reserva.AnioInicio = reservaDto.AnioInicio
	reserva.AnioFinal = reservaDto.AnioFinal
	reserva.MesInicio = reservaDto.MesInicio
	reserva.MesFinal = reservaDto.MesFinal
	reserva.DiaInicio = reservaDto.DiaInicio
	reserva.DiaFinal = reservaDto.DiaFinal
	reserva.Dias = reservaDto.Dias

	reserva.HotelID = reservaDto.HotelID
	reserva.ClienteID = reservaDto.ClienteID

	reserva.Hotel = hotel
	reserva.Cliente = cliente

	reserva = reservaClient.InsertReserva(reserva)

	reservaDto.ID = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetReservasById(id int) (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.GetReservasById(id)
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		if reserva.ClienteID == 0 {
			return reservasDto, e.NewBadRequestApiError("Reservas No Encontradas")
		}

		reservaDto.ID = reserva.ID
		reservaDto.HotelID = reserva.Hotel.ID
		reservaDto.ClienteID = reserva.Cliente.ID
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

func (s *reservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva = reservaClient.GetReservaById(id)
	var reservaDto dto.ReservaDto

	if reserva.ID == 0 {
		return reservaDto, e.NewBadRequestApiError("Reserva No Encontrada")
	}

	reservaDto.ID = reserva.ID
	reservaDto.HotelID = reserva.Hotel.ID
	reservaDto.ClienteID = reserva.Cliente.ID
	reservaDto.AnioInicio = reserva.AnioInicio
	reservaDto.AnioFinal = reserva.AnioFinal
	reservaDto.MesInicio = reserva.MesInicio
	reservaDto.MesFinal = reserva.MesFinal
	reservaDto.DiaInicio = reserva.DiaInicio
	reservaDto.DiaFinal = reserva.DiaFinal
	reservaDto.Dias = reserva.Dias

	return reservaDto, nil
}

func (s *reservaService) GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int) {
	
	var reservas model.Reservas = reservaClient.GetDisponibilidad(id)
	var hotel model.Hotel = hotelClient.GetHotelById(id)
	
	disponibilidad = hotel.Cant_Hab

	for _, reserva := range reservas {
		if reserva.HotelID == id && reserva.AnioFinal >= AnioInicio && reserva.AnioInicio <= AnioFinal && reserva.MesFinal >= MesInicio && reserva.MesInicio <= MesFinal && reserva.DiaFinal >= DiaInicio && reserva.DiaInicio <= DiaFinal {
			disponibilidad --
		}
	}

	return disponibilidad
}

func (s *reservaService) GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, e.ApiError) {
	
	var reservas model.Reservas = reservaClient.GetReservasByDate()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		if reserva.AnioFinal >= AnioInicio && reserva.AnioInicio <= AnioFinal && reserva.MesFinal >= MesInicio && reserva.MesInicio <= MesFinal && reserva.DiaFinal >= DiaInicio && reserva.DiaInicio <= DiaFinal {
			reservaDto.ID = reserva.ID
			reservaDto.HotelID = reserva.Hotel.ID
			reservaDto.ClienteID = reserva.Cliente.ID
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
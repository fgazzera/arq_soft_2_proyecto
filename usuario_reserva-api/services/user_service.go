package services

import (
	clienteClient "backend/clients/cliente"
	hotelClient "backend/clients/hotel"
	reservaClient "backend/clients/reserva"
	imagenClient "backend/clients/imagen"

	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type clienteService struct{}

type clienteServiceInterface interface {
	GetClienteById(id int) (dto.ClienteDto, e.ApiError)
	GetClienteByUsername(username string) (dto.ClienteDto, e.ApiError)
	GetClienteByEmail(email string) (dto.ClienteDto, e.ApiError)
	InsertCliente(clienteDto dto.ClienteDto) (dto.ClienteDto, e.ApiError)
	GetHoteles() (dto.HotelesDto, e.ApiError)
	GetImagenesByHotelId(hotelID int) (dto.ImagenesDto, e.ApiError)
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetReservasById(id int) (dto.ReservasDto, e.ApiError)
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int)
	GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, e.ApiError)
}

var (
	ClienteService clienteServiceInterface
)

func init() {
	ClienteService = &clienteService{}
}

func (s *clienteService) GetClienteById(id int) (dto.ClienteDto, e.ApiError) {

	var cliente model.Cliente = clienteClient.GetClienteById(id)
	var clienteDto dto.ClienteDto

	if cliente.ID == 0 {
		return clienteDto, e.NewBadRequestApiError("Cliente No Encontrado")
	}

	clienteDto.ID = cliente.ID
	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Password = cliente.Password
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *clienteService) GetClienteByUsername(username string) (dto.ClienteDto, e.ApiError) {
	var cliente model.Cliente = clienteClient.GetClienteByUsername(username)
	var clienteDto dto.ClienteDto

	if cliente.UserName == "" {
		return clienteDto, e.NewBadRequestApiError("Cliente No Encontrado")
	}

	clienteDto.ID = cliente.ID
	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Password = cliente.Password
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *clienteService) GetClienteByEmail(email string) (dto.ClienteDto, e.ApiError) {
	var cliente model.Cliente = clienteClient.GetClienteByEmail(email)
	var clienteDto dto.ClienteDto

	if cliente.Email == "" {
		return clienteDto, e.NewBadRequestApiError("Cliente No Encontrado")
	}

	clienteDto.ID = cliente.ID
	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Password = cliente.Password
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *clienteService) InsertCliente(clienteDto dto.ClienteDto) (dto.ClienteDto, e.ApiError) {

	var cliente model.Cliente

	cliente.Name = clienteDto.Name
	cliente.LastName = clienteDto.LastName
	cliente.UserName = clienteDto.UserName
	cliente.Password = clienteDto.Password
	cliente.Email = clienteDto.Email

	cliente = clienteClient.InsertCliente(cliente)

	clienteDto.ID = cliente.ID

	return clienteDto, nil
}

func (s *clienteService) GetHoteles() (dto.HotelesDto, e.ApiError) {

	var hoteles model.Hoteles = hotelClient.GetHoteles()
	var hotelesDto dto.HotelesDto

	for _, hotel := range hoteles {
		var hotelDto dto.HotelDto
		hotelDto.ID = hotel.ID
		hotelDto.Nombre = hotel.Nombre
		hotelDto.Descripcion = hotel.Descripcion
		hotelDto.Email = hotel.Email
		hotelDto.Cant_Hab = hotel.Cant_Hab
		hotelDto.Amenities = hotel.Amenities

		hotelesDto = append(hotelesDto, hotelDto)
	}

	return hotelesDto, nil
}

func (i *clienteService) GetImagenesByHotelId(hotelID int) (dto.ImagenesDto, e.ApiError) {
	var imagenes model.Imagenes = imagenClient.GetImagenesByHotelId(hotelID)
	var imagenesDto dto.ImagenesDto

	for _, imagen := range imagenes {
		var imagenDto dto.ImagenDto

		if imagen.HotelID == 0 {
			return imagenesDto, e.NewBadRequestApiError("Imagenes No Encontradas")
		}

		imagenDto.ID = imagen.ID
		imagenDto.Url = imagen.Url
		imagenDto.HotelID = imagen.Hotel.ID

		imagenesDto = append(imagenesDto, imagenDto)
	}

	return imagenesDto, nil
}

func (s *clienteService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.ID == 0 {
		return hotelDto, e.NewBadRequestApiError("Hotel No Encontrado")
	}

	hotelDto.ID = hotel.ID
	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Email = hotel.Email
	hotelDto.Cant_Hab = hotel.Cant_Hab
	hotelDto.Amenities = hotel.Amenities

	return hotelDto, nil
}

func (s *clienteService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

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

func (s *clienteService) GetReservasById(id int) (dto.ReservasDto, e.ApiError) {

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

func (s *clienteService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

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

func (s *clienteService) GetDisponibilidad(id, AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (disponibilidad int) {
	
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

func (s *clienteService) GetReservasByDate(AnioInicio, AnioFinal, MesInicio, MesFinal, DiaInicio, DiaFinal int) (dto.ReservasDto, e.ApiError) {
	
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
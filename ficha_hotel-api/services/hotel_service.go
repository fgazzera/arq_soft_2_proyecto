package services

import (
	hotelDao "ficha_hotel-api/daos/hotel"
	"ficha_hotel-api/dtos"
	"ficha_hotel-api/model"
	e "ficha_hotel-api/utils/errors"

	"github.com/streadway/amqp"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id string) (dtos.HotelDto, e.ApiError)
	InsertHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
	UpdateHotelById(id string, hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

const rabbitMQURL = "amqp://user:password@localhost:5672"

func init() {
	HotelService = &hotelService{}
}

func setupRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

func (s *hotelService) GetHotelById(id string) (dtos.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelDao.GetHotelById(id)
	var hotelDto dtos.HotelDto

	if hotel.ID.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotelDto.ID = hotel.ID.Hex()
	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Email = hotel.Email
	hotelDto.Ciudad = hotel.Ciudad
	hotelDto.Images = hotel.Images
	hotelDto.CantHab = hotel.CantHab
	hotelDto.Amenities = hotel.Amenities

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Nombre = hotelDto.Nombre
	hotel.Descripcion = hotelDto.Descripcion
	hotel.Email = hotelDto.Email
	hotel.Ciudad = hotelDto.Ciudad
	hotel.Images = hotelDto.Images
	hotel.CantHab = hotelDto.CantHab
	hotel.Amenities = hotelDto.Amenities

	hotel = hotelDao.InsertHotel(hotel)

	if hotel.ID.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("error in insert")
	}

	hotelDto.ID = hotel.ID.Hex()

	// Enviar un mensaje a RabbitMQ después de crear el hotel
	conn, ch, err := setupRabbitMQ()
	if err != nil {
		return hotelDto, e.NewInternalServerApiError("Error al configurar RabbitMQ", err)
	}
	defer conn.Close()
	defer ch.Close()

	exchangeName := "hotel_insert" // Nombre del intercambio en RabbitMQ
	routingKey := "hotel.created"  // Clave de enrutamiento para la creación de hoteles

	message := "Se ha creado un nuevo hotel con ID: " + hotel.ID.Hex() // Mensaje a enviar

	err = ch.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		return hotelDto, e.NewInternalServerApiError("Error al declarar el intercambio en RabbitMQ", err)
	}

	err = ch.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		return hotelDto, e.NewInternalServerApiError("Error al publicar el mensaje en RabbitMQ", err)
	}

	return hotelDto, nil
}

func (s *hotelService) UpdateHotelById(id string, hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelDao.GetHotelById(id)

	if hotel.ID.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotel.Nombre = hotelDto.Nombre
	hotel.Descripcion = hotelDto.Descripcion
	hotel.Email = hotelDto.Email
	hotel.Ciudad = hotelDto.Ciudad
	hotel.Images = hotelDto.Images
	hotel.CantHab = hotelDto.CantHab
	hotel.Amenities = hotelDto.Amenities

	hotel = hotelDao.UpdateHotel(hotel)
	hotelDto.ID = hotel.ID.Hex()

	// Enviar un mensaje a RabbitMQ después de actualizar el hotel
	conn, ch, err := setupRabbitMQ()
	if err != nil {
		return hotelDto, e.NewInternalServerApiError("Error al configurar RabbitMQ", err)
	}
	defer conn.Close()
	defer ch.Close()

	exchangeName := "hotel_update" // Nombre del intercambio en RabbitMQ
	routingKey := "hotel.updated"  // Clave de enrutamiento para la actualización de hoteles

	message := "Se ha actualizado un hotel con ID: " + hotel.ID.Hex() // Mensaje a enviar

	err = ch.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		return hotelDto, e.NewInternalServerApiError("Error al declarar el intercambio en RabbitMQ", err)
	}

	err = ch.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		return hotelDto, e.NewInternalServerApiError("Error al publicar el mensaje en RabbitMQ", err)
	}

	return hotelDto, nil
}

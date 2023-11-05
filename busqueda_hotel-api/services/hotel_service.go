package services

import (
	hotelDao "busqueda_hotel-api/daos/hotel" // Asegúrate de importar el paquete DAO correcto
	"busqueda_hotel-api/dtos"
	model "busqueda_hotel-api/models"
	e "busqueda_hotel-api/utils/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotel(id string) (dtos.HotelDto, e.ApiError)
	CreateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
	UpdateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
	GetAllHotels() (dtos.HotelsDto, e.ApiError)
	GetHotelsByCiudad(ciudad string) (dtos.HotelsDto, e.ApiError)
	GetDisponibilidad(ciudad string, fechainicio string, fechafinal string) (dtos.HotelsDisponibilidadDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetAllHotels() (dtos.HotelsDto, e.ApiError) {

	var hotelDtos dtos.HotelsDto
	hotelDtos.Hotels = []dtos.HotelDto{}
	hotelClient := hotelDao.NewHotelSolrDAO()
	hotels, err := hotelClient.GetAll()
	if err != nil {
		return hotelDtos, e.NewBadRequestApiError("error al obtener hoteles")
	}

	// Mapear los modelos de hoteles a DTOs

	for _, hotel := range hotels {
		hotelDto := dtos.HotelDto{
			ID:          hotel.ID,
			Nombre:      hotel.Nombre,
			Descripcion: hotel.Descripcion,
			Email:       hotel.Email,
			Ciudad:      hotel.Ciudad,
			Images:      hotel.Images,
			CantHab:     hotel.CantHab,
			Amenities:   hotel.Amenities,
		}
		hotelDtos.Hotels = append(hotelDtos.Hotels, hotelDto)
	}

	return hotelDtos, nil
}

func (s *hotelService) GetHotelsByCiudad(ciudad string) (dtos.HotelsDto, e.ApiError) {

	var hotelDtos dtos.HotelsDto
	hotelDtos.Hotels = []dtos.HotelDto{}
	hotelClient := hotelDao.NewHotelSolrDAO()
	hotels, err := hotelClient.GetByCiudad(ciudad)
	if err != nil {
		return hotelDtos, e.NewBadRequestApiError("error al obtener hoteles")
	}

	// Mapear los modelos de hoteles a DTOs

	for _, hotel := range hotels {
		hotelDto := dtos.HotelDto{
			ID:          hotel.ID,
			Nombre:      hotel.Nombre,
			Descripcion: hotel.Descripcion,
			Email:       hotel.Email,
			Ciudad:      hotel.Ciudad,
			Images:      hotel.Images,
			CantHab:     hotel.CantHab,
			Amenities:   hotel.Amenities,
		}
		hotelDtos.Hotels = append(hotelDtos.Hotels, hotelDto)
	}

	return hotelDtos, nil
}

type DisponibilidadResult struct {
	HotelID        string
	Disponibilidad bool
}

func (s *hotelService) GetDisponibilidad(ciudad string, fechainicio string, fechafinal string) (dtos.HotelsDisponibilidadDto, e.ApiError) {
	var hotelDtos dtos.HotelsDisponibilidadDto
	hotelDtos.Hotels = []dtos.HotelDisponibilidadDto{}
	hotelClient := hotelDao.NewHotelSolrDAO()
	var hotels []*model.Hotel
	var err error
	if ciudad == "" {
		hotels, err = hotelClient.GetAll()
	} else {
		hotels, err = hotelClient.GetByCiudad(ciudad)
	}

	if err != nil {
		return hotelDtos, e.NewBadRequestApiError("error al obtener hoteles")
	}

	// Crear un canal para recibir resultados de disponibilidad
	disponibilidadCh := make(chan DisponibilidadResult, len(hotels))

	// Crear una WaitGroup para esperar que todas las goroutines terminen
	var wg sync.WaitGroup

	// Mapear los modelos de hoteles a DTOs y consultar la disponibilidad concurrentemente
	for _, hotel := range hotels {

		hotelDto := dtos.HotelDisponibilidadDto{
			ID:          hotel.ID,
			Nombre:      hotel.Nombre,
			Descripcion: hotel.Descripcion,
			Email:       hotel.Email,
			Ciudad:      hotel.Ciudad,
			Images:      hotel.Images,
			CantHab:     hotel.CantHab,
			Amenities:   hotel.Amenities,
		}

		// Incrementar el contador de WaitGroup para cada goroutine
		wg.Add(1)

		go func(hotel *model.Hotel, hotelDto dtos.HotelDisponibilidadDto) {
			defer wg.Done() // Decrementar el contador cuando la goroutine termine

			// Realizar la solicitud de disponibilidad al servicio
			disponibilidad, err := checkDisponibilidad(hotel.ID, fechainicio, fechafinal)
			if err != nil {
				println("Error al hacer el get a user-res-api: ", err.Error())
				disponibilidadCh <- DisponibilidadResult{HotelID: hotel.ID, Disponibilidad: false}
				return
			}

			// Enviar el resultado al canal
			disponibilidadCh <- DisponibilidadResult{HotelID: hotel.ID, Disponibilidad: disponibilidad}
		}(hotel, hotelDto)

		// Agregar el hotelDto a la lista de DTOs
		hotelDtos.Hotels = append(hotelDtos.Hotels, hotelDto)
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	// Cerrar el canal después de que todas las goroutines hayan enviado sus resultados
	close(disponibilidadCh)

	// Crear un mapa para almacenar los resultados de disponibilidad
	disponibilidadMap := make(map[string]bool)

	// Recopilar resultados de disponibilidad y agregarlos a los DTOs de hoteles
	for result := range disponibilidadCh {
		disponibilidadMap[result.HotelID] = result.Disponibilidad
	}

	// Asignar los valores de disponibilidad desde el mapa a los DTOs de hoteles
	for i, hotel := range hotelDtos.Hotels {
		disponibilidad := disponibilidadMap[hotel.ID]
		println("Hotel: ", hotel.ID, " Disponibilidad: ", disponibilidad)
		hotelDtos.Hotels[i].Disponibilidad = disponibilidad
	}

	return hotelDtos, nil
}

func checkDisponibilidad(hotelID string, fechainicio string, fechafinal string) (bool, error) {
	// Construye la URL de la solicitud de disponibilidad
	url := fmt.Sprintf("http://user-res-api:8002/hotel/%s/disponibilidad?fecha-inicio=%s&fecha-final=%s", hotelID, fechainicio, fechafinal)

	// Realiza la solicitud HTTP GET al servicio de disponibilidad
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Verifica si la respuesta fue exitosa (código 200)
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("La solicitud de disponibilidad no fue exitosa. Código de respuesta: %d", resp.StatusCode)
	}

	// Lee el cuerpo de la respuesta HTTP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Decodifica la respuesta JSON
	var disponibilidadResponse struct {
		Disponibilidad bool `json:"disponibilidad"`
	}
	if err := json.Unmarshal(body, &disponibilidadResponse); err != nil {
		return false, err
	}

	// Retorna la disponibilidad obtenida de la respuesta
	return disponibilidadResponse.Disponibilidad, nil
}

func (s *hotelService) GetHotel(id string) (dtos.HotelDto, e.ApiError) {

	var hotelDto dtos.HotelDto
	hotelClient := hotelDao.NewHotelSolrDAO()
	hotel, err := hotelClient.Get(id)
	if err != nil {
		return hotelDto, e.NewBadRequestApiError(err.Error())
	}

	if hotel.ID == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotelDto.ID = hotel.ID
	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Email = hotel.Email
	hotelDto.Ciudad = hotel.Ciudad
	hotelDto.Images = hotel.Images
	hotelDto.CantHab = hotel.CantHab
	hotelDto.Amenities = hotel.Amenities

	return hotelDto, nil
}

func (s *hotelService) CreateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {
	var hotel model.Hotel

	// Aquí deberías mapear los campos del DTO al modelo Hotel
	hotel.ID = hotelDto.ID
	hotel.Nombre = hotelDto.Nombre
	hotel.Descripcion = hotelDto.Descripcion
	hotel.Email = hotelDto.Email
	hotel.Ciudad = hotelDto.Ciudad
	hotel.Images = hotelDto.Images
	hotel.CantHab = hotelDto.CantHab
	hotel.Amenities = hotelDto.Amenities

	hotelClient := hotelDao.NewHotelSolrDAO()
	err := hotelClient.Create(&hotel)

	if err != nil {
		return hotelDto, e.NewBadRequestApiError(err.Error())
	}
	hotelDto.ID = hotel.ID

	return hotelDto, nil
}

func (s *hotelService) UpdateHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {
	var hotel model.Hotel

	// Aquí deberías mapear los campos del DTO al modelo Hotel
	hotel.ID = hotelDto.ID
	hotel.Nombre = hotelDto.Nombre
	hotel.Descripcion = hotelDto.Descripcion
	hotel.Email = hotelDto.Email
	hotel.Ciudad = hotelDto.Ciudad
	hotel.Images = hotelDto.Images
	hotel.CantHab = hotelDto.CantHab
	hotel.Amenities = hotelDto.Amenities

	hotelClient := hotelDao.NewHotelSolrDAO()
	err := hotelClient.Update(&hotel)

	if err != nil {
		return hotelDto, e.NewBadRequestApiError("error in update")
	}
	hotelDto.ID = hotel.ID
	hotelDto.ID = hotel.ID

	return hotelDto, nil
}

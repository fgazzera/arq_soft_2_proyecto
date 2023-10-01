package services

import (
	"ficha_hotel-api/cache"
	hotelDao "ficha_hotel-api/daos/hotel"
	"ficha_hotel-api/dtos"
	model "ficha_hotel-api/model"
	e "ficha_hotel-api/utils/errors"
	"fmt"
	"time"

	json "github.com/json-iterator/go"
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

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id string) (dtos.HotelDto, e.ApiError) {

	time.Sleep(15 * time.Second)

	// get from cache
	var cacheDTO dtos.HotelDto
	cacheBytes := cache.Get(id)
	if cacheBytes != nil {
		fmt.Println("Found in cache!")
		_ = json.Unmarshal(cacheBytes, &cacheDTO)
		return cacheDTO, nil
	}

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
	hotelDto.Cant_Hab = hotel.Cant_Hab
	hotelDto.Amenities = hotel.Amenities

	// save in cache
	hotelBytes, _ := json.Marshal(hotelDto)
	cache.Set(id, hotelBytes)
	fmt.Println("Saved in cache!")

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Nombre = hotelDto.Nombre
	hotel.Descripcion = hotelDto.Descripcion
	hotel.Email = hotelDto.Email
	hotel.Ciudad = hotelDto.Ciudad
	hotel.Images = hotelDto.Images
	hotel.Cant_Hab = hotelDto.Cant_Hab
	hotel.Amenities = hotelDto.Amenities

	hotel = hotelDao.InsertHotel(hotel)

	if hotel.ID.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("error in insert")
	}

	hotelDto.ID = hotel.ID.Hex()
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
	hotel.Cant_Hab = hotelDto.Cant_Hab
	hotel.Amenities = hotelDto.Amenities

	hotel = hotelDao.Update(hotel)
	hotelDto.ID = hotel.ID.Hex()

	// save in cache
	hotelBytes, _ := json.Marshal(hotelDto)
	cache.Set(id, hotelBytes)
	fmt.Println("Saved in cache!")

	return hotelDto, nil
}

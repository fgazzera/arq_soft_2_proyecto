package services

import (
	userDao "usuario_reserva-api/daos/user"

	dtos "usuario_reserva-api/dtos"
	models "usuario_reserva-api/models"
	e "usuario_reserva-api/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	InsertUser(userDto dtos.UserDto) (dtos.UserDto, e.ApiError)
	GetUserById(id int) (dtos.UserDto, e.ApiError)
	GetUserByUsername(username string) (dtos.UserDto, e.ApiError)
	GetUserByEmail(email string) (dtos.UserDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) InsertUser(userDto dtos.UserDto) (dtos.UserDto, e.ApiError) {

	var user models.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Password = userDto.Password
	user.Email = userDto.Email

	user = userDao.InsertUser(user)

	userDto.ID = user.ID

	return userDto, nil
}

func (s *userService) GetUserById(id int) (dtos.UserDto, e.ApiError) {

	var user models.User = userDao.GetUserById(id)
	var userDto dtos.UserDto

	if user.ID == 0 {
		return userDto, e.NewBadRequestApiError("Usuario No Encontrado")
	}

	userDto.ID = user.ID
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Email = user.Email

	return userDto, nil
}

func (s *userService) GetUserByUsername(username string) (dtos.UserDto, e.ApiError) {
	var user models.User = userDao.GetUserByUsername(username)
	var userDto dtos.UserDto

	if user.UserName == "" {
		return userDto, e.NewBadRequestApiError("Usuario No Encontrado")
	}

	userDto.ID = user.ID
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Email = user.Email

	return userDto, nil
}

func (s *userService) GetUserByEmail(email string) (dtos.UserDto, e.ApiError) {
	var user models.User = userDao.GetUserByEmail(email)
	var userDto dtos.UserDto

	if user.Email == "" {
		return userDto, e.NewBadRequestApiError("Usuario No Encontrado")
	}

	userDto.ID = user.ID
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Email = user.Email

	return userDto, nil
}

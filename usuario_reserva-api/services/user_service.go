package services

import (
	"fmt"
	//"time"
	"usuario_reserva-api/cache"
	userDao "usuario_reserva-api/daos/user"
	"usuario_reserva-api/dtos"
	"usuario_reserva-api/models"
	e "usuario_reserva-api/utils/errors"

	json "github.com/json-iterator/go"
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

	// save in cache
	userBytes, _ := json.Marshal(userDto)
	cache.Set(fmt.Sprintf("user_%d", userDto.ID), userBytes)
	cache.Set(fmt.Sprintf("user_username_%d", userDto.UserName), userBytes)
	cache.Set(fmt.Sprintf("user_email_%d", userDto.Email), userBytes)
	fmt.Println("Saved user in cache!")

	return userDto, nil
}

func (s *userService) GetUserById(id int) (dtos.UserDto, e.ApiError) {

	//time.Sleep(15 * time.Second)

	// Genera una clave de caché única para el usuario
	cacheKey := fmt.Sprintf("user_%d", id)

	// get from cache
	var cacheDTO dtos.UserDto
	cacheBytes := cache.Get(cacheKey)
	if cacheBytes != nil {
		fmt.Println("Found user in cache!")
		_ = json.Unmarshal(cacheBytes, &cacheDTO)
		return cacheDTO, nil
	}

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

	// save in cache
	userBytes, _ := json.Marshal(userDto)
	cache.Set(cacheKey, userBytes)
	fmt.Println("Saved user in cache!")

	return userDto, nil
}

func (s *userService) GetUserByUsername(username string) (dtos.UserDto, e.ApiError) {

	//time.Sleep(15 * time.Second)

	// Genera una clave de caché única para el usuario
	cacheKey := fmt.Sprintf("user_username_%s", username)

	// get from cache
	var cacheDTO dtos.UserDto
	cacheBytes := cache.Get(cacheKey)
	if cacheBytes != nil {
		fmt.Println("Found user in cache!")
		_ = json.Unmarshal(cacheBytes, &cacheDTO)
		return cacheDTO, nil
	}

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

	// save in cache
	userBytes, _ := json.Marshal(userDto)
	cache.Set(cacheKey, userBytes)
	fmt.Println("Saved user in cache!")

	return userDto, nil
}

func (s *userService) GetUserByEmail(email string) (dtos.UserDto, e.ApiError) {

	//time.Sleep(15 * time.Second)

	// Genera una clave de caché única para el usuario
	cacheKey := fmt.Sprintf("user_email_%s", email)

	// get from cache
	var cacheDTO dtos.UserDto
	cacheBytes := cache.Get(cacheKey)
	if cacheBytes != nil {
		fmt.Println("Found user in cache!")
		_ = json.Unmarshal(cacheBytes, &cacheDTO)
		return cacheDTO, nil
	}

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

	// save in cache
	userBytes, _ := json.Marshal(userDto)
	cache.Set(cacheKey, userBytes)
	fmt.Println("Saved user in cache!")

	return userDto, nil
}

// Genera una clave de caché única para usuarios
func generateUserCacheKey(id int) string {
	return fmt.Sprintf("user:%d", id)
}

package userController

import (
	"net/http"
	"strconv"
	"time"
	dtos "usuario_reserva-api/dtos"
	service "usuario_reserva-api/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertUser(c *gin.Context) {
	var userDto dtos.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}

func GetUserById(c *gin.Context) {
	log.Debug("ID de usuario para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dtos.UserDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func GetUserByUsername(c *gin.Context) {
	log.Debug("Usuario a cargar: " + c.Param("username"))

	username := c.Param("username")
	var userDto dtos.UserDto

	userDto, err := service.UserService.GetUserByUsername(username)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func GetUserByEmail(c *gin.Context) {
	log.Debug("Usuario a cargar: " + c.Param("email"))

	email := c.Param("email")
	var userDto dtos.UserDto

	userDto, err := service.UserService.GetUserByEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	token := generateToken(userDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	response := struct {
		Token   string       `json:"token"`
		Usuario dtos.UserDto `json:"usuario"`
	}{
		Token:   token,
		Usuario: userDto,
	}

	c.JSON(http.StatusOK, response)
}

func generateToken(loginDto dtos.UserDto) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = loginDto.ID
	claims["expiration"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return ""
	}

	return tokenString
}

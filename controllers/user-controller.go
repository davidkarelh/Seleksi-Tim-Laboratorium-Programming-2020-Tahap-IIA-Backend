package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/dto"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/helper"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserController is a ....
type IUserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type UserController struct {
	userService services.IUserService
	jwtService  services.IJWTService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userService services.IUserService, jwtService services.IJWTService) IUserController {
	return &UserController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *UserController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *UserController) Profile(context *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := context.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := c.jwtService.ValidateToken(tokenString)

	if err != nil {
		res := helper.BuildErrorResponse("OK", "Validate Token Error", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := fmt.Sprintf("%v", claims["username"])
	user := c.userService.Profile(username)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}

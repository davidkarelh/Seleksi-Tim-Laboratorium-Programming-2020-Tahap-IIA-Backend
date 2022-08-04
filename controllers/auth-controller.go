package controllers

import (
	"net/http"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/dto"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/helper"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/services"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type AuthController struct {
	jwtService  services.IJWTService
	userService services.IUserService
	service     services.IAuthService
}

func NewAuthController(jwtService services.IJWTService, userService services.IUserService, service services.IAuthService) IAuthController {
	return &AuthController{
		jwtService:  jwtService,
		userService: userService,
		service:     service,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.service.VerifyCredential(loginDTO.UserName, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.Username)
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Username or/and password invalid", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if c.service.IsDuplicateUsername(registerDTO.UserName) {
		response := helper.BuildErrorResponse("Failed to process request", "Username already exists", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err := c.service.AddToRegisterVerification(registerDTO)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildResponse(true, "You have been registered. Please wait for the verification.", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
	return
}

func (c *AuthController) Profile(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.service.VerifyCredential(loginDTO.UserName, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.Username)
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Username or/and password invalid", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

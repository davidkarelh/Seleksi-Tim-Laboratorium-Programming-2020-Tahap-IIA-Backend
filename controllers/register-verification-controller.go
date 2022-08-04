package controllers

import (
	"net/http"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/helper"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/services"
	"github.com/gin-gonic/gin"
)

type IRegisterVerificationController interface {
	GetAllRegisterVerifications(*gin.Context)
}

type RegisterVerificationController struct {
	service services.IRegisterVerificationService
}

func NewRegisterVerificationController(service services.IRegisterVerificationService) IRegisterVerificationController {
	return &RegisterVerificationController{
		service: service,
	}
}

// func (c * RegisterVerificationController) {}
func (c *RegisterVerificationController) GetAllRegisterVerifications(ctx *gin.Context) {
	users := c.service.GetAllRegisterVerifications()
	res := helper.BuildResponse(true, "List of all register verifications.", users)
	ctx.JSON(http.StatusOK, res)
}

package controllers

import (
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/services"
)

type IHistoryController interface {
}

type HistoryController struct {
	service services.IHistoryService
}

func NewHistoryController(service services.IHistoryService) IHistoryController {
	return &HistoryController{
		service: service,
	}
}

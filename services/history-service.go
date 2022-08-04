package services

import (
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/repository"
)

type IHistoryService interface {
}
type HistoryService struct {
	repository repository.IHistoryRepository
}

func NewHistoryService(repo repository.IHistoryRepository) IHistoryService {
	return &HistoryService{
		repository: repo,
	}
}

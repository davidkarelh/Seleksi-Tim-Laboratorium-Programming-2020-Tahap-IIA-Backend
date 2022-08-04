package services

import (
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/repository"
)

type IRegisterVerificationService interface {
	GetAllRegisterVerifications() []entity.RegisterVerification
}

type RegisterVerificationService struct {
	repository repository.IRegisterVerificationRepository
}

func NewRegisterVerificationService(repository repository.IRegisterVerificationRepository) IRegisterVerificationService {
	return &RegisterVerificationService{
		repository: repository,
	}
}

// func (s *RegisterVerificationService) {}
func (s *RegisterVerificationService) GetAllRegisterVerifications() []entity.RegisterVerification {
	var user entity.RegisterVerification
	return s.repository.Find(user)
}

package services

import (
	"log"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/dto"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/repository"
	"github.com/mashingan/smapping"
)

type IUserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	VerifyCredential(username string, passwrod string) interface{}
}

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) VerifyCredential(username string, password string) interface{} {
	return s.repository.VerifyCredential(username, password)
}

func (s *UserService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := s.repository.UpdateUser(userToUpdate)
	return updatedUser
}

func (s *UserService) Profile(username string) entity.User {
	return s.repository.Find(entity.User{Username: username})[0]
}

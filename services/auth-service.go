package services

import (
	"fmt"
	"log"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/dto"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type IAuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	IsDuplicateUsername(username string) bool
	AddToRegisterVerification(user dto.RegisterDTO) error
}

type AuthService struct {
	userRepository                 repository.IUserRepository
	registerVerificationRepository repository.IRegisterVerificationRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.IUserRepository, registerVerificationRepository repository.IRegisterVerificationRepository) IAuthService {
	return &AuthService{
		userRepository:                 userRep,
		registerVerificationRepository: registerVerificationRepository,
	}
}

func (service *AuthService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		fmt.Println(comparedPassword)
		fmt.Println(username)
		fmt.Println(v.Username)
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *AuthService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *AuthService) IsDuplicateUsername(username string) bool {
	res1, _ := service.userRepository.IsDuplicateUsername(username)
	res2 := service.registerVerificationRepository.IsDuplicateUsername(username)

	// fmt.Println(user.Username)
	// fmt.Println(username)
	// fmt.Println(user.Username == username)
	// fmt.Println(res.Error != nil)
	// fmt.Println(res.Error)
	return (res1.Error == nil || res2.Error == nil)
}

func (service *AuthService) AddToRegisterVerification(user dto.RegisterDTO) error {
	userToCreate := entity.RegisterVerification{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	errSave := service.registerVerificationRepository.Insert(userToCreate)
	return errSave

}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

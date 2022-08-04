package repository

import (
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/helper"
	"gorm.io/gorm"
)

type IRegisterVerificationRepository interface {
	Insert(user entity.RegisterVerification) error
	IsDuplicateUsername(username string) *gorm.DB
	Find(user entity.RegisterVerification) []entity.RegisterVerification
}

type RegisterVerificationConnection struct {
	connection *gorm.DB
}

func NewRegisterVerificationRepository(connection *gorm.DB) IRegisterVerificationRepository {
	return &RegisterVerificationConnection{
		connection: connection,
	}
}

func (db *RegisterVerificationConnection) Insert(user entity.RegisterVerification) error {
	user.Password = helper.HashAndSalt([]byte(user.Password))
	saveResult := db.connection.Save(&user)
	return saveResult.Error
}

func (db *RegisterVerificationConnection) IsDuplicateUsername(username string) *gorm.DB {
	var user entity.RegisterVerification
	return db.connection.Where("user_name = ?", username).First(&user)
}

func (db *RegisterVerificationConnection) Find(user entity.RegisterVerification) []entity.RegisterVerification {
	var users []entity.RegisterVerification
	db.connection.Where(&user).Find(&users)
	return users
}

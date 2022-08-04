package repository

import (
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/helper"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(username string, pasword string) interface{}
	IsDuplicateUsername(username string) (*gorm.DB, entity.User)
	FindByUsername(username string) entity.User
	Find(user entity.User) []entity.User
}

type UserConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserConnection{
		connection: db,
	}
}

func (db *UserConnection) InsertUser(user entity.User) entity.User {
	user.Password = helper.HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *UserConnection) UpdateUser(user entity.User) entity.User {
	user.Password = helper.HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *UserConnection) VerifyCredential(username string, pasword string) interface{} {
	var user entity.User
	res := db.connection.Where(&entity.User{Username: username}).First(&user)

	if res.Error == nil {
		return user
	}

	return nil
}

func (db *UserConnection) IsDuplicateUsername(username string) (*gorm.DB, entity.User) {
	var user entity.User
	x := db.connection.Where("user_name = ?", username).First(&user)
	return x, user
}

func (db *UserConnection) FindByUsername(username string) entity.User {
	var user entity.User
	db.connection.Where("username = ?", username).Take(&user)
	return user
}

func (db *UserConnection) Find(user entity.User) []entity.User {
	var users []entity.User
	db.connection.Where(&user).Find(&users)
	return users
}

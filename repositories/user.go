package repositories

import (
	"literature/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	// err := r.db.Preload("Literatur.User").Preload("Collection.User").Find(&users).Error

	return users, err
}

// func (r *repository) GetUser(ID int) (models.User, error) {
// 	var user models.User
// 	err := r.db.First(&user, ID).Error

// 	return user, err
// }

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Literatur.User").Preload("Collection.User").First(&user, ID).Error

	return user, err
}

// func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
// 	err := r.db.Save(&user).Error

// 	return user, err
// }

// func (r *repository) CreateUser(user models.User) (models.User, error) {
// 	// err := r.db.Create(&user).Error
// 	// err := r.db.Preload("Journey").Preload("Journey.User").Preload("Bookmark.Journey.User").Preload("Bookmark.User").Create(&user).Error
// 	err := r.db.Preload("Literatur").Preload("Literatur.User").Preload("Collection.Literatur.User").Preload("Collection.User").Create(&user).Error

// 	return user, err
// }

// // func (r *repository) GetUser(ID int) (models.User, error) {
// // 	var user models.User
// // 	err := r.db.First(&user, ID).Error

// // 	return user, err
// // }

// func (r *repository) FindUsers() ([]models.User, error) {
// 	var users []models.User
// 	err := r.db.Preload("Literatur").Preload("Literatur.User").Preload("Collection.Literatur.User").Preload("Collection.User").Find(&users).Error

// 	return users, err
// }

// func (r *repository) GetUser(ID int) (models.User, error) {
// 	var user models.User
// 	err := r.db.Preload("Literatur").Preload("Literatur.User").Preload("Collection.Literatur.User").Preload("Collection.User").First(&user, ID).Error

// 	return user, err
// }

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

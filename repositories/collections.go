package repositories

import (
	"literature/models"

	"gorm.io/gorm"
)

type CollectionRepository interface {
	CreateCollection(User models.Collection) (models.Collection, error)
	FindCollection() ([]models.Collection, error)
	GetCollection(ID int) (models.Collection, error)
	UpdateCollection(User models.Collection) (models.Collection, error)
	DeleteCollection(User models.Collection) (models.Collection, error)
	GetCollectionByUserID(userID int) ([]models.Collection, error)
	GetCollectionByLiteratur(Lit int) ([]models.Collection, error)
}

func RepositoryCollection(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCollection(collection models.Collection) (models.Collection, error) {
	err := r.db.Preload("User").Preload("Literatur.User").Create(&collection).Error
	return collection, err
}

func (r *repository) FindCollection() ([]models.Collection, error) {
	var collection []models.Collection
	err := r.db.Preload("User").Preload("Literatur.User").Find(&collection).Error
	return collection, err
}

func (r *repository) GetCollection(ID int) (models.Collection, error) {
	var collection models.Collection
	err := r.db.Preload("User").Preload("Literatur.User").First(&collection, ID).Error
	// err := r.db.Preload("User").Preload("Literatur.User").Where("user_id = ?", User ID).Find(&collection).Error

	return collection, err
}

func (r *repository) UpdateCollection(collection models.Collection) (models.Collection, error) {
	err := r.db.Preload("User").Preload("Literatur.User").Save(&collection).Error

	return collection, err
}

func (r *repository) DeleteCollection(collection models.Collection) (models.Collection, error) {
	err := r.db.Preload("User").Preload("Literatur.User").Delete(&collection).Error

	return collection, err
}

func (r *repository) GetCollectionByUserID(userID int) ([]models.Collection, error) {
	var collections []models.Collection
	err := r.db.Debug().Preload("User").Preload("Literatur").Where("user_id= ?", userID).Find(&collections).Error

	return collections, err
}

func (r *repository) GetCollectionByLiteratur(Lit int) ([]models.Collection, error) {
	var collection []models.Collection
	err := r.db.Where("literatur_id=?", Lit).Preload("User").Preload("Literatur").Preload("Literatur.User").Find(&collection).Error
	return collection, err
}

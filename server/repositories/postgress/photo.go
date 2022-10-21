package postgress

import (
	"finalproject/server/models"
	"finalproject/server/repositories"

	"gorm.io/gorm"
)

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) repositories.PhotoRepository {
	return &photoRepo{
		db: db,
	}
}

func (a *photoRepo) PhotoAdd(photo *models.Photo) error {

	err := a.db.Create(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *photoRepo) PhotoGet() (*[]models.Photo, error) {
	var photo []models.Photo

	err := a.db.Find(&photo).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (a *photoRepo) GetUserInfo(id int) *models.User {
	var user models.User

	err := a.db.Where("Id = ?", id).Find(&user).Error
	if err != nil {
		return nil
	}

	return &user
}

func (a *photoRepo) PhotoUpdate(photo *models.Photo) (*models.Photo, error) {

	err := a.db.Model(&photo).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoUrl: photo.PhotoUrl, UpdatedAt: photo.UpdatedAt}).Error
	if err != nil {
		return nil, err
	}

	var dataPhoto models.Photo

	err = a.db.Where("Id = ?", photo.Id).Find(&dataPhoto).Error
	if err != nil {
		return nil, err
	}
	return &dataPhoto, nil
}

func (a *photoRepo) PhotoDelete(id int) error {
	err := a.db.Delete(&models.Photo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *photoRepo) PhotoAuth(userId, photoId int) bool {

	var photo models.Photo

	err := a.db.First(&photo).Error
	if err != nil {
		return false
	}

	if photo.UserId == userId {
		return true
	}

	return false
}

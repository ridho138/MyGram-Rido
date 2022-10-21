package postgress

import (
	"finalproject/server/models"
	"finalproject/server/repositories"

	"gorm.io/gorm"
)

type socmedRepo struct {
	db *gorm.DB
}

func NewSocmedRepo(db *gorm.DB) repositories.SocialMediaRepository {
	return &socmedRepo{
		db: db,
	}
}

func (a *socmedRepo) SocialMediaAdd(socmed *models.SocialMedia) error {

	err := a.db.Create(&socmed).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *socmedRepo) SocialMediaGet() (*[]models.SocialMedia, error) {
	var socmed []models.SocialMedia

	err := a.db.Find(&socmed).Error
	if err != nil {
		return nil, err
	}
	return &socmed, nil
}

func (a *socmedRepo) GetSocmedUserInfo(id int) *models.User {
	var user models.User

	err := a.db.Where("Id = ?", id).Find(&user).Error
	if err != nil {
		return nil
	}

	return &user
}

func (a *socmedRepo) SocialMediaUpdate(socmed *models.SocialMedia) (*models.SocialMedia, error) {

	err := a.db.Model(&socmed).Updates(models.SocialMedia{Name: socmed.Name, SocialMediaUrl: socmed.SocialMediaUrl}).Error
	if err != nil {
		return nil, err
	}

	var dataSocmed models.SocialMedia

	err = a.db.Where("Id = ?", socmed.Id).Find(&dataSocmed).Error
	if err != nil {
		return nil, err
	}
	return &dataSocmed, nil
}

func (a *socmedRepo) SocialMediaDelete(id int) error {
	err := a.db.Delete(&models.SocialMedia{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *socmedRepo) SocmedAuth(userId, SocmedId int) bool {

	var socmed models.SocialMedia

	err := a.db.First(&socmed).Error
	if err != nil {
		return false
	}

	if socmed.UserId == userId {
		return true
	}

	return false
}

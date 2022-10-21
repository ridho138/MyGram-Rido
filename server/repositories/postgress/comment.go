package postgress

import (
	"finalproject/server/models"
	"finalproject/server/repositories"

	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) repositories.CommentRepository {
	return &commentRepo{
		db: db,
	}
}

func (a *commentRepo) CommentAdd(comment *models.Comment) error {

	err := a.db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *commentRepo) CommentGet() (*[]models.Comment, error) {
	var comment []models.Comment

	err := a.db.Find(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (a *commentRepo) GetCommentUserInfo(id int) *models.User {
	var user models.User

	err := a.db.Where("Id = ?", id).Find(&user).Error
	if err != nil {
		return nil
	}

	return &user
}

func (a *commentRepo) GetCommentPhotoInfo(id int) *models.Photo {
	var photo models.Photo

	err := a.db.Where("Id = ?", id).Find(&photo).Error
	if err != nil {
		return nil
	}

	return &photo
}

func (a *commentRepo) CommentUpdate(comment *models.Comment) (*models.Comment, error) {

	err := a.db.Model(&comment).Updates(models.Comment{Message: comment.Message}).Error
	if err != nil {
		return nil, err
	}

	var dataComment models.Comment

	err = a.db.Where("Id = ?", comment.Id).Find(&dataComment).Error
	if err != nil {
		return nil, err
	}
	return &dataComment, nil
}

func (a *commentRepo) CommentDelete(id int) error {
	err := a.db.Delete(&models.Comment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *commentRepo) CommentAuth(userId, commentId int) bool {

	var comment models.Comment

	err := a.db.First(&comment).Error
	if err != nil {
		return false
	}

	if comment.UserId == userId {
		return true
	}

	return false
}

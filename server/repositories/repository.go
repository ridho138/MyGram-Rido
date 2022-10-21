package repositories

import "finalproject/server/models"

type UserRepository interface {
	UserRegister(user *models.User) error
	UserLogin(email string) *models.User
	UserUpdate(user *models.User) (*models.User, error)
	UserDelete(id int) error
}

type PhotoRepository interface {
	PhotoAdd(photo *models.Photo) error
	PhotoGet() (*[]models.Photo, error)
	GetUserInfo(id int) *models.User
	PhotoUpdate(photo *models.Photo) (*models.Photo, error)
	PhotoDelete(id int) error
	PhotoAuth(userId, photoId int) bool
}

type CommentRepository interface {
	CommentAdd(comment *models.Comment) error
	CommentGet() (*[]models.Comment, error)
	CommentUpdate(comment *models.Comment) (*models.Comment, error)
	CommentDelete(id int) error
	GetCommentUserInfo(id int) *models.User
	CommentAuth(userId, photoId int) bool
	GetCommentPhotoInfo(id int) *models.Photo
}

type SocialMediaRepository interface {
	SocialMediaAdd(socmed *models.SocialMedia) error
	SocialMediaGet() (*[]models.SocialMedia, error)
	SocialMediaUpdate(socmed *models.SocialMedia) (*models.SocialMedia, error)
	SocialMediaDelete(id int) error
	GetSocmedUserInfo(id int) *models.User
	SocmedAuth(userId, photoId int) bool
}

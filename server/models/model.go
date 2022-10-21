package models

import "time"

type User struct {
	Id        int    `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Photo struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	Id        int `gorm:"primaryKey"`
	UserId    int
	PhotoId   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SocialMedia struct {
	Id             int `gorm:"primaryKey"`
	Name           string
	SocialMediaUrl string
	UserId         int
}

type DataUserReq struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
	Age      int    `validate:"required,gte=9"`
}

type DataUserLoginReq struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

type DataUserUpdate struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
}

type DataPhotoReq struct {
	Title    string `validate:"required"`
	Caption  string
	PhotoUrl string `validate:"required"`
}

type DataSocMedReq struct {
	Name           string `validate:"required"`
	SocialMediaUrl string `validate:"required"`
}

type DataCommentReq struct {
	Message string `validate:"required"`
	PhotoId int
}

type PhotoResult struct {
	Id        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      DataUserUpdate
}

type CommentResult struct {
	Id        int
	UserId    int
	PhotoId   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      DataUserUpdate
	Photo     Photo
}

type SocmedResult struct {
	Id             int
	Name           string
	SocialMediaUrl string
	UserId         int
	User           DataUserUpdate
}

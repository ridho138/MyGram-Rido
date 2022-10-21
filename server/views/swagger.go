package views

import "finalproject/server/models"

type SwaggerResponse struct {
	Status  string
	Message string
	Error   string
}

type SwaggerUserRegister struct {
	SwaggerResponse
	Payload models.User
}

type SwaggerUserLogin struct {
	SwaggerResponse
	Token string
}

type SwaggerPhotoAdd struct {
	SwaggerResponse
	Payload models.Photo
}

type SwaggerPhotoGet struct {
	SwaggerResponse
	Payload []models.PhotoResult
}

type SwaggerCommentAdd struct {
	SwaggerResponse
	Payload models.Comment
}

type SwaggerCommentGet struct {
	SwaggerResponse
	Payload []models.CommentResult
}

type SwaggerSocmedAdd struct {
	SwaggerResponse
	Payload models.SocialMedia
}

type SwaggerSocmedGet struct {
	SwaggerResponse
	Payload []models.SocmedResult
}

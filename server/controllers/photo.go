package controllers

import (
	"finalproject/server/models"
	"finalproject/server/repositories"
	"finalproject/server/views"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoController struct {
	photoRepo repositories.PhotoRepository
}

func NewPhotoController(photoRepo repositories.PhotoRepository) *PhotoController {
	validate = validator.New()
	return &PhotoController{
		photoRepo: photoRepo,
	}
}

// PhotoAdd godoc
// @Summary Add new photo
// @Decription Add new photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param data body models.DataPhotoReq true "Add New Photo"
// @Success 200 {object} views.SwaggerPhotoAdd
// @Router /photos [post]
func (a *PhotoController) PhotoAdd(ctx *gin.Context) {
	var dataReq models.DataPhotoReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_ADD_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_ADD_FAILED",
			Error:   message,
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))

	var dataPhoto = models.Photo{
		Title:     dataReq.Title,
		Caption:   dataReq.Caption,
		PhotoUrl:  dataReq.PhotoUrl,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = a.photoRepo.PhotoAdd(&dataPhoto)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_ADD_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "PHOTO_ADD_SUCCESS",
		Payload: dataPhoto,
	})
}

// PhotoGet godoc
// @Summary Get all photo
// @Decription Get all photo
// @Tags Photo
// @Accept json
// @Produce json
// @Success 200 {object} views.SwaggerPhotoGet
// @Router /photos [get]
func (a *PhotoController) PhotoGet(ctx *gin.Context) {

	photos, err := a.photoRepo.PhotoGet()
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_ADD_FAILED",
			Error:   err.Error(),
		})
		return
	}

	var result []models.PhotoResult

	for _, x := range *photos {
		dataUser := a.photoRepo.GetUserInfo(x.UserId)
		result = append(result,
			models.PhotoResult{
				Id:        x.Id,
				Title:     x.Title,
				Caption:   x.Caption,
				PhotoUrl:  x.PhotoUrl,
				UserId:    x.UserId,
				CreatedAt: x.CreatedAt,
				UpdatedAt: x.UpdatedAt,
				User: models.DataUserUpdate{
					Username: dataUser.Username,
					Email:    dataUser.Email,
				},
			},
		)
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "PHOTO_ADD_SUCCESS",
		Payload: result,
	})
}

// PhotoUpdate godoc
// @Summary Update photo
// @Decription Update photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param photoId path int true "Photo ID"
// @Param data body models.DataPhotoReq true "Update Photo"
// @Success 200 {object} views.SwaggerPhotoAdd
// @Router /photos/{photoId} [put]
func (a *PhotoController) PhotoUpdate(ctx *gin.Context) {
	var dataReq models.DataPhotoReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_UPDATE_FAILED",
			Error:   message,
		})
		return
	}

	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}
	userId, _ := strconv.Atoi(ctx.GetString("id"))
	auth := PhotoAuth(a, userId, photoId)
	if !auth {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_UPDATE_FAILED",
			Error:   "You dont have access",
		})
		return
	}

	var dataUpdate = models.Photo{
		Id:        photoId,
		Title:     dataReq.Title,
		Caption:   dataReq.Caption,
		PhotoUrl:  dataReq.PhotoUrl,
		UpdatedAt: time.Now(),
	}

	dataPhoto, err := a.photoRepo.PhotoUpdate(&dataUpdate)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "PHOTO_UPDATE_SUCCESS",
		Payload: dataPhoto,
	})
}

// PhotoDelete godoc
// @Summary Update photo
// @Decription Update photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param photoId path int true "Photo ID"
// @Success 200 {object} views.SwaggerResponse
// @Router /photos/{photoId} [delete]
func (a *PhotoController) PhotoDelete(ctx *gin.Context) {

	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_DELETE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))
	auth := PhotoAuth(a, userId, photoId)
	if !auth {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_UPDATE_FAILED",
			Error:   "You dont have access",
		})
		return
	}

	err = a.photoRepo.PhotoDelete(photoId)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "PHOTO_DELETE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "Your photo has been successfully deleted",
	})
}

func PhotoAuth(a *PhotoController, userId, photoId int) bool {

	auth := a.photoRepo.PhotoAuth(userId, photoId)

	return auth

}

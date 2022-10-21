package controllers

import (
	"finalproject/server/models"
	"finalproject/server/repositories"
	"finalproject/server/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SocmedController struct {
	socmedRepo repositories.SocialMediaRepository
}

func NewSocmedController(socmedRepo repositories.SocialMediaRepository) *SocmedController {
	validate = validator.New()
	return &SocmedController{
		socmedRepo: socmedRepo,
	}
}

// SocmedAdd godoc
// @Summary Add new social media
// @Decription Add new social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Param data body models.DataSocMedReq true "Add New social media"
// @Success 200 {object} views.SwaggerSocmedAdd
// @Router /socialmedias [post]
func (a *SocmedController) SocmedAdd(ctx *gin.Context) {
	var dataReq models.DataSocMedReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_ADD_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_ADD_FAILED",
			Error:   message,
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))

	var dataSocmed = models.SocialMedia{
		Name:           dataReq.Name,
		SocialMediaUrl: dataReq.SocialMediaUrl,
		UserId:         userId,
	}

	err = a.socmedRepo.SocialMediaAdd(&dataSocmed)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_ADD_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "SOCIAL_MEDIA_ADD_SUCCESS",
		Payload: dataSocmed,
	})
}

// SocmedGet godoc
// @Summary Get social media
// @Decription Get social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Success 200 {object} views.SwaggerSocmedGet
// @Router /socialmedias [get]
func (a *SocmedController) SocmedGet(ctx *gin.Context) {

	socmeds, err := a.socmedRepo.SocialMediaGet()
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_GET_FAILED",
			Error:   err.Error(),
		})
		return
	}

	var result []models.SocmedResult

	for _, x := range *socmeds {
		dataUser := a.socmedRepo.GetSocmedUserInfo(x.UserId)
		result = append(result,
			models.SocmedResult{
				Id:             x.Id,
				Name:           x.Name,
				SocialMediaUrl: x.SocialMediaUrl,
				UserId:         x.UserId,
				User: models.DataUserUpdate{
					Username: dataUser.Username,
					Email:    dataUser.Email,
				},
			},
		)
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "SOCIAL_MEDIA_GET_SUCCESS",
		Payload: result,
	})
}

// SocmedUpdate godoc
// @Summary Update social media
// @Decription Update social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Param socmedId path int true "Social Media ID"
// @Param data body models.DataSocMedReq true "Update social media"
// @Success 200 {object} views.SwaggerSocmedAdd
// @Router /socialmedias/{socmedId} [put]
func (a *SocmedController) SocmedUpdate(ctx *gin.Context) {
	var dataReq models.DataSocMedReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_UPDATE_FAILED",
			Error:   message,
		})
		return
	}

	socmedId, err := strconv.Atoi(ctx.Param("socmedId"))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))
	auth := SocmedAuth(a, userId, socmedId)
	if !auth {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_UPDATE_FAILED",
			Error:   "You dont have access",
		})
		return
	}

	var dataUpdate = models.SocialMedia{
		Name:           dataReq.Name,
		SocialMediaUrl: dataReq.SocialMediaUrl,
		UserId:         userId,
	}

	dataSocmed, err := a.socmedRepo.SocialMediaUpdate(&dataUpdate)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "SOCIAL_MEDIA_UPDATE_SUCCESS",
		Payload: dataSocmed,
	})
}

// SocmedDelete godoc
// @Summary Delete social media
// @Decription Delete social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Param socmedId path int true "Social Media ID"
// @Success 200 {object} views.SwaggerResponse
// @Router /socialmedias/{socmedId} [delete]
func (a *SocmedController) SocmedDelete(ctx *gin.Context) {

	socmedId, err := strconv.Atoi(ctx.Param("socmedId"))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_DELETE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))
	auth := SocmedAuth(a, userId, socmedId)
	if !auth {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_UPDATE_FAILED",
			Error:   "You dont have access",
		})
		return
	}

	err = a.socmedRepo.SocialMediaDelete(socmedId)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "SOCIAL_MEDIA_DELETE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "Your social media has been successfully deleted",
	})
}

func SocmedAuth(a *SocmedController, userId, socmedId int) bool {

	auth := a.socmedRepo.SocmedAuth(userId, socmedId)

	return auth

}

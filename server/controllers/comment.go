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

type CommentController struct {
	commentRepo repositories.CommentRepository
}

func NewCommentController(commentRepo repositories.CommentRepository) *CommentController {
	validate = validator.New()
	return &CommentController{
		commentRepo: commentRepo,
	}
}

// CommentAdd godoc
// @Summary Add new comment
// @Decription Add new comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param data body models.DataCommentReq true "Add New comment"
// @Success 200 {object} views.SwaggerCommentAdd
// @Router /comments [post]
func (a *CommentController) CommentAdd(ctx *gin.Context) {
	var dataReq models.DataCommentReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_ADD_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_ADD_FAILED",
			Error:   message,
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))

	var dataComment = models.Comment{
		UserId:    userId,
		PhotoId:   dataReq.PhotoId,
		Message:   dataReq.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = a.commentRepo.CommentAdd(&dataComment)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_ADD_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "COMMENT_ADD_SUCCESS",
		Payload: dataComment,
	})
}

// CommentGet godoc
// @Summary get comment
// @Decription get comment
// @Tags Comment
// @Accept json
// @Produce json
// @Success 200 {object} views.SwaggerCommentGet
// @Router /comments [get]
func (a *CommentController) CommentGet(ctx *gin.Context) {

	comments, err := a.commentRepo.CommentGet()
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_GET_FAILED",
			Error:   err.Error(),
		})
		return
	}

	var result []models.CommentResult

	for _, x := range *comments {
		dataPhoto := a.commentRepo.GetCommentPhotoInfo(x.PhotoId)
		dataUser := a.commentRepo.GetCommentUserInfo(dataPhoto.UserId)
		result = append(result, models.CommentResult{
			Id:        x.Id,
			UserId:    x.UserId,
			PhotoId:   x.PhotoId,
			Message:   x.Message,
			CreatedAt: x.CreatedAt,
			UpdatedAt: x.UpdatedAt,
			User: models.DataUserUpdate{
				Username: dataUser.Username,
				Email:    dataUser.Email,
			},
			Photo: *dataPhoto,
		})
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "COMMENT_GET_SUCCESS",
		Payload: result,
	})
}

// CommentUpdate godoc
// @Summary Update comment
// @Decription Update comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Param data body models.DataCommentReq true "Update comment"
// @Success 200 {object} views.SwaggerCommentAdd
// @Router /comments/{commentId} [put]
func (a *CommentController) CommentUpdate(ctx *gin.Context) {
	var dataReq models.DataCommentReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_UPDATE_FAILED",
			Error:   message,
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))
	auth := CommentAuth(a, userId, commentId)
	if !auth {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_UPDATE_FAILED",
			Error:   "You dont have access",
		})
		return
	}

	var dataUpdate = models.Comment{
		Message:   dataReq.Message,
		UpdatedAt: time.Now(),
	}

	dataComment, err := a.commentRepo.CommentUpdate(&dataUpdate)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "COMMENT_UPDATE_SUCCESS",
		Payload: dataComment,
	})
}

// CommentDelete godoc
// @Summary Delete comment
// @Decription Delete comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} views.SwaggerResponse
// @Router /comments/{commentId} [delete]
func (a *CommentController) CommentDelete(ctx *gin.Context) {

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_DELETE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	userId, _ := strconv.Atoi(ctx.GetString("id"))
	auth := CommentAuth(a, userId, commentId)
	if !auth {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_UPDATE_FAILED",
			Error:   "You dont have access",
		})
		return
	}

	err = a.commentRepo.CommentDelete(commentId)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "COMMENT_DELETE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully deleted",
	})
}

func CommentAuth(a *CommentController, userId, commentId int) bool {

	auth := a.commentRepo.CommentAuth(userId, commentId)

	return auth

}

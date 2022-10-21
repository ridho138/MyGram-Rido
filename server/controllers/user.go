package controllers

import (
	"finalproject/helper"
	"finalproject/server/models"
	"finalproject/server/repositories"
	"finalproject/server/views"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

type UserController struct {
	actionRepo repositories.UserRepository
}

func NewUserController(actionRepo repositories.UserRepository) *UserController {
	validate = validator.New()
	return &UserController{
		actionRepo: actionRepo,
	}
}

// UserRegister godoc
// @Summary Add new user
// @Decription Add new user
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.DataUserReq true "Add New User"
// @Success 200 {object} views.SwaggerUserRegister
// @Router /users/register [post]
func (a *UserController) UserRegister(ctx *gin.Context) {
	var dataReq models.DataUserReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_REGISTER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_REGISTER_FAILED",
			Error:   message,
		})
		return
	}

	pwdEncrypt, err := bcrypt.GenerateFromPassword([]byte(dataReq.Password), bcrypt.DefaultCost)

	var dataUser = models.User{
		Username:  dataReq.Username,
		Email:     dataReq.Email,
		Password:  string(pwdEncrypt),
		Age:       dataReq.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = a.actionRepo.UserRegister(&dataUser)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_REGISTER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "USER_REGISTER_SUCCESS",
		Payload: dataUser,
	})
}

// UserLogin godoc
// @Summary User Login
// @Decription User Login
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.DataUserLoginReq true "User Login"
// @Success 200 {object} views.SwaggerUserLogin
// @Router /users/login [post]
func (a *UserController) UserLogin(ctx *gin.Context) {
	var dataReq models.DataUserLoginReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_LOGIN_FAILED",
			Error:   err.Error(),
		})
		return
	}

	user := a.actionRepo.UserLogin(dataReq.Email)
	if user == nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusNotFound,
			Message: "USER_LOGIN_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataReq.Password))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusUnauthorized,
			Message: "USER_LOGIN_FAILED",
			Error:   err.Error(),
		})
		return
	}
	fmt.Println("User Login ", user)
	token, err := helper.GenerateToken(user.Email, strconv.Itoa(user.Id))
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_LOGIN_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "USER_LOGIN_SUCCESS",
		Token:   token,
	})
}

// UserUpdate godoc
// @Summary User Update
// @Decription User Update
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.DataUserUpdate true "User Update"
// @Success 200 {object} views.SwaggerUserRegister
// @Router /users [put]
func (a *UserController) UserUpdate(ctx *gin.Context) {
	var dataReq models.DataUserUpdate
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = validate.Struct(dataReq)
	if err != nil {
		message := FieldValidation(err)
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_UPDATE_FAILED",
			Error:   message,
		})
		return
	}

	id, _ := strconv.Atoi(ctx.GetString("id"))

	var dataUpdate = models.User{
		Id:        id,
		Username:  dataReq.Username,
		Email:     dataReq.Email,
		UpdatedAt: time.Now(),
	}

	dataUser, err := a.actionRepo.UserUpdate(&dataUpdate)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_UPDATE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "USER_UPDATE_SUCCESS",
		Payload: dataUser,
	})
}

// UserDelete godoc
// @Summary User Delete
// @Decription User Delete
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} views.SwaggerResponse
// @Router /users [delete]
func (a *UserController) UserDelete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("id"))

	err := a.actionRepo.UserDelete(id)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "USER_DELETE_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "Your account has been successfully deleted",
	})
}

func FieldValidation(err error) []string {
	var message []string
	for _, err := range err.(validator.ValidationErrors) {

		switch err.Field() {
		case "Email":
			message = append(message, "Format email salah")
		case "Password":
			message = append(message, "Panjang password minimal 6 karakter")
		case "Username":
			message = append(message, "Username tidak boleh kosong")
		case "Age":
			message = append(message, "Minimal umur 9")
		case "Title":
			message = append(message, "Title tidak boleh kosong")
		case "PhotoUrl":
			message = append(message, "Title tidak boleh kosong")
		case "Message":
			message = append(message, "Message tidak boleh kosong")
		case "Name":
			message = append(message, "Name tidak boleh kosong")
		case "SocialMediaUrl":
			message = append(message, "Social Media URL tidak boleh kosong")
		}
	}
	return message
}

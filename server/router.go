package server

import (
	"finalproject/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	user    *controllers.UserController
	photo   *controllers.PhotoController
	comment *controllers.CommentController
	socmed  *controllers.SocmedController
}

func NewRouter(router *gin.Engine, user *controllers.UserController, photo *controllers.PhotoController, comment *controllers.CommentController, socmed *controllers.SocmedController) *Router {
	return &Router{
		router:  router,
		user:    user,
		photo:   photo,
		comment: comment,
		socmed:  socmed,
	}
}

func (r *Router) Start(port string) {

	r.router.POST("/users/register", r.user.UserRegister)
	r.router.POST("/users/login", r.user.UserLogin)
	r.router.PUT("/users", CheckAuth, r.user.UserUpdate)
	r.router.DELETE("/users", CheckAuth, r.user.UserDelete)

	r.router.POST("/photos", CheckAuth, r.photo.PhotoAdd)
	r.router.GET("/photos", CheckAuth, r.photo.PhotoGet)
	r.router.PUT("/photos/:photoId", CheckAuth, r.photo.PhotoUpdate)
	r.router.DELETE("/photos/:photoId", CheckAuth, r.photo.PhotoDelete)

	r.router.POST("/comments", CheckAuth, r.comment.CommentAdd)
	r.router.GET("/comments", CheckAuth, r.comment.CommentGet)
	r.router.PUT("/comments/:commentId", CheckAuth, r.comment.CommentUpdate)
	r.router.DELETE("/comments/:commentId", CheckAuth, r.comment.CommentDelete)

	r.router.POST("/socialmedias", CheckAuth, r.socmed.SocmedAdd)
	r.router.GET("/socialmedias", CheckAuth, r.socmed.SocmedGet)
	r.router.PUT("/socialmedias/:socmedId", CheckAuth, r.socmed.SocmedUpdate)
	r.router.DELETE("/socialmedias/:socmedId", CheckAuth, r.socmed.SocmedDelete)

	r.router.Run(port)
}

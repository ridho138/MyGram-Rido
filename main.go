package main

import (
	"finalproject/db"
	"finalproject/server"
	"finalproject/server/controllers"
	"finalproject/server/repositories/postgress"

	"github.com/gin-gonic/gin"
)

// @title MyGram API
// @description Share and comment ur photo
// @version v1.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:4000
// @contact.name Teguh Ridho Afdilla
// @contact.email teguh.afdilla138@gmail.com
func main() {
	db := db.ConnectGorm()

	userRepo := postgress.NewUserRepo(db)
	userHandler := controllers.NewUserController(userRepo)

	photoRepo := postgress.NewPhotoRepo(db)
	photoHandler := controllers.NewPhotoController(photoRepo)

	commentRepo := postgress.NewCommentRepo(db)
	commentHandler := controllers.NewCommentController(commentRepo)

	socmedRepo := postgress.NewSocmedRepo(db)
	socmedHandler := controllers.NewSocmedController(socmedRepo)

	router := gin.Default()
	server.NewRouter(router, userHandler, photoHandler, commentHandler, socmedHandler).Start(":4000")
}

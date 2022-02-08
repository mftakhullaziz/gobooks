package main

import (
	"github.com/amifth/ApiGo/configuration"
	"github.com/amifth/ApiGo/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = configuration.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer configuration.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group(("api/auth"))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

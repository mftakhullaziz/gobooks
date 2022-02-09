package main

import (
	"github.com/amifth/ApiGo/configuration"
	"github.com/amifth/ApiGo/controller"
	"github.com/amifth/ApiGo/middleware"
	"github.com/amifth/ApiGo/repository"
	"github.com/amifth/ApiGo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = configuration.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer configuration.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group(("api/auth"), middleware.AuthorizeJWT(jwtService))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

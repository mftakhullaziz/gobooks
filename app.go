package main

import (
	"fmt"

	"github.com/amifth/gorest/configuration"
	"github.com/amifth/gorest/controller"
	_ "github.com/amifth/gorest/docs"
	"github.com/amifth/gorest/middleware"
	"github.com/amifth/gorest/repository"
	"github.com/amifth/gorest/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	db             = configuration.SetupDatabaseConnection()
	userRepository = repository.NewUserRepository(db)
	bookRepository = repository.NewBookRepository(db)
	jwtService     = service.NewJWTService()
	userService    = service.NewUserService(userRepository)
	authService    = service.NewAuthService(userRepository)
	bookService    = service.NewBookService(bookRepository)
	authController = controller.NewAuthController(authService, jwtService)
	userController = controller.NewUserController(userService, jwtService)
	bookController = controller.NewBookController(bookService, jwtService)
)

// @title           Apigo - Spec Documentation API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	defer configuration.CloseDatabaseConnection(db)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/login", authController.Login)
			authRoutes.POST("/register", authController.Register)
		}

		dataUserRoutes := v1.Group("/user")
		{
			dataUserRoutes.GET("/all", userController.AllUser)
		}

		userRoutes := v1.Group("/user", middleware.AuthorizeJWT(jwtService))
		{
			userRoutes.GET("/profile/:id", userController.Profile)
			userRoutes.PUT("/profile", userController.Update)
		}

		bookRoutes := v1.Group("/books", middleware.AuthorizeJWT(jwtService))
		{
			bookRoutes.GET("/", bookController.All)
			bookRoutes.POST("/", bookController.Insert)
			bookRoutes.GET("/:id", bookController.FindByID)
			bookRoutes.PUT("/:id", bookController.Update)
			bookRoutes.DELETE("/:id", bookController.Delete)
		}
	}

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	fmt.Println("Documentation API : http://localhost:8080/swagger/index.html")
	err := r.Run()
	if err != nil {
		return
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

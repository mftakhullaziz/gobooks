package main

import (
	"fmt"

	"github.com/amifth/apigo-gin/configuration"
	"github.com/amifth/apigo-gin/controller"
	_ "github.com/amifth/apigo-gin/docs"
	"github.com/amifth/apigo-gin/middleware"
	"github.com/amifth/apigo-gin/repository"
	"github.com/amifth/apigo-gin/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = configuration.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
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
			userRoutes.GET("/profile", userController.Profile)
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

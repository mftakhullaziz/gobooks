package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mftakhullaziz/gorest/dto"
	"github.com/mftakhullaziz/gorest/entity"
	"github.com/mftakhullaziz/gorest/helper"
	"github.com/mftakhullaziz/gorest/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Login Authentication godoc
// @Summary      login to account
// @Description  login to your account
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        email    query     string  false  "email"  Format(email)
// @Param        password    query     string  false  "password"  Format(password)
// @Success      200  {object}  map[string]interface{}
// @Router       /auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(http.StatusCreated, true, _apiRequest, false, v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check your credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Register Authentication godoc
// @Summary      register to account
// @Description  register
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "name"  Format(name)
// @Param        email    query     string  false  "email"  Format(email)
// @Param        password    query     string  true  "password"  Format(password)
// @Success      200  {object}  map[string]interface{}
// @Router       /auth/register [post]
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	// Check if duplicate email for register
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token
		if token != "" {
			response := helper.BuildResponse(http.StatusCreated, true, _apiCreated, false, createUser)
			ctx.JSON(http.StatusCreated, response)
		}
	}
}

package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amifth/gorest/dto"
	"github.com/amifth/gorest/entity"
	"github.com/amifth/gorest/helper"
	"github.com/amifth/gorest/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
	return &bookController{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
}

/*All Book godoc
@Summary      book data
@Description  fetch all book data
@Tags         book
@Accept       json
@Produce      json
@Param 		  Authorization header string true "Bearer"
@Success      200  {object}  map[string]interface{}
@Router       /all [get]*/
func (c *bookController) All(context *gin.Context) {
	var books = c.bookService.All()
	res := helper.BuildResponse("200", true, "Successful!", books)
	context.JSON(http.StatusOK, res)
}

/*FindByID Book godoc
@Summary      book data
@Description  get book data by id
@Tags         book
@Accept       json
@Produce      json
@Param 		  Authorization header string true "Bearer"
@Param        bookId    query     string  false  "bookId"  Format(bookId)
@Success      200  {object}  map[string]interface{}
@Router       /book/:id [get]*/
func (c *bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var books = c.bookService.FindByID(id)
	if (books == entity.Book{}) {
		res := helper.BuildErrorResponse("Data not found", "no data given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse("200", true, "Successful!", books)
		context.JSON(http.StatusOK, res)
	}
}

/*Insert User godoc
@Summary      book data
@Description  get book data
@Tags         book
@Accept       json
@Produce      json
@Param 		  Authorization header string true "Bearer"
@Param        bookId    query     string  false  "bookId"  Format(bookId)
@Success      200  {object}  map[string]interface{}
@Router       /insert/:id [put]*/
func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Request failed", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertUserID, err := strconv.ParseUint(userID, 10, 64)
		// fmt.Println(convertUserID)
		if err == nil {
			bookCreateDTO.UserID = convertUserID
		}
		result := c.bookService.Insert(bookCreateDTO)
		response := helper.BuildResponse("200", true, "Successful!", result)
		context.JSON(http.StatusCreated, response)
	}
}

/*Update Book godoc
@Summary      user account
@Description  user update
@Tags         user
@Accept       json
@Produce      json
@Param 		 Authorization header string true "Bearer"
@Success      200  {object}  map[string]interface{}
@Router       /user/update [put]*/
func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Request failed", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdateDTO)
		response := helper.BuildResponse("200", true, "Successful!", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

/*Delete Book godoc
@Summary      book data
@Description  delete book data by id
@Tags         book
@Accept       json
@Produce      json
@Param 		  Authorization header string true "Bearer"
@Success      200  {object}  map[string]interface{}
@Router       /books/:id [delete]*/
func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed get id", "No Param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	book.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		res := helper.BuildResponse("200", true, "Delete Successful!", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "Permission denied", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

/*getUserIDByToken Book godoc
@Summary      book data
@Description  fetch token book data
@Tags         book
@Accept       json
@Produce      json
@Param 		  Authorization header string true "Bearer"
@Success      200  {object}  map[string]interface{}
@Router       /all [get]*/
func (c *bookController) getUserIDByToken(token string) string {
	uToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := uToken.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

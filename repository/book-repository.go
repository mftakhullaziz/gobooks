package repository

import (
	"github.com/amifth/ApiGo/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbConnection *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConnection,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("user").Find(&b)
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("user").Find(&b)
	return b
}

package service

import (
	bookModel "clean-arsitektur/internal/model/book"
	"database/sql"
)

type bookService struct {
	db *sql.DB
}

type BookService interface {
	AddBookService(bookModel.BookRequest) (string, error)
}

func NewBookService(db *sql.DB) BookService {
	return &bookService{db}
}

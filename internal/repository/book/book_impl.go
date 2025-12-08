package repository

import (
	bookModel "clean-arsitektur/internal/model/book"
	"database/sql"
)

type bookRepository struct {
	db *sql.DB
}

type BookRepository interface {
	AddBook(bookModel.BookRequest) error
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db}
}

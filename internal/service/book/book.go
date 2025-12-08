package service

import (
	bookModel "clean-arsitektur/internal/model/book"
	bookRepo "clean-arsitektur/internal/repository/book"
)

func (q *bookService) AddBookService(req bookModel.BookRequest) (string, error) {

	if err := bookRepo.NewBookRepository(q.db).AddBook(req); err != nil {
		return "failed", err
	}

	return "success", nil
}

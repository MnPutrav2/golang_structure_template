package repository

import bookModel "clean-arsitektur/internal/model/book"

func (q *bookRepository) AddBook(req bookModel.BookRequest) error {

	if _, err := q.db.Exec("INSERT INTO book(id, name) VALUES($1, $2)", req.ID, req.Name); err != nil {
		return err
	}

	return nil
}

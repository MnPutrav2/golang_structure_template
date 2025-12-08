package route

import (
	handlerBook "clean-arsitektur/internal/handler/book"
	"database/sql"
	"net/http"
)

func BookRoute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "POST":
			handlerBook.BookCreate(db)(w, r)
		}

	}
}

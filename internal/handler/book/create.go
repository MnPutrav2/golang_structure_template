package handler

import (
	"clean-arsitektur/internal/model"
	bookModel "clean-arsitektur/internal/model/book"
	bookService "clean-arsitektur/internal/service/book"
	logging "clean-arsitektur/pkg/logging"
	"clean-arsitektur/pkg/middleware"
	"clean-arsitektur/pkg/util"
	"database/sql"
	"encoding/json"
	"net/http"
)

func BookCreate(db *sql.DB) http.HandlerFunc {
	return middleware.CORS(
		middleware.RateLimiter(1, 1, func(w http.ResponseWriter, r *http.Request) {
			body, err := util.BodyDecoder[bookModel.BookRequest](r)
			if err != nil {
				res, _ := json.Marshal(model.ResponseMessage{Status: "failed", Message: "failed decode body"})
				logging.Log(err.Error(), "ERROR", r)
				w.WriteHeader(500)
				w.Write(res)

				return
			}

			mess, err := bookService.NewBookService(db).AddBookService(body)
			if err != nil {
				res, _ := json.Marshal(model.ResponseMessage{Status: "failed", Message: mess})
				logging.Log(mess, "ERROR", r)
				w.WriteHeader(400)
				w.Write(res)

				return
			}

			res, _ := json.Marshal(model.ResponseMessage{Status: "success", Message: mess})
			logging.Log(mess, "INFO", r)
			w.WriteHeader(200)
			w.Write(res)
		}),
	)
}

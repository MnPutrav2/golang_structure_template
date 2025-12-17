package make

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Template(dir, name, ty string) {
	module := moduleReader()

	hdlTemp := fmt.Sprintf(`
		package handler

		import (
			"%s/internal/model"
			logging "%s/pkg/logging"
			"%s/pkg/middleware"
			"database/sql"
			"encoding/json"
			"net/http"
		)

		func %sController(db *sql.DB) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {

				switch r.Method {
				case "GET":
					// Write method get
					// Example
					handlerName(db)(w, r)
				case "POST":
					// Write method post
				case "PUT":
					// Write method put
				case "DELETE":
					// Write method delete
				}

			}
		}

		func handlerName(db *sql.DB) http.HandlerFunc {
			return middleware.CORS(
				middleware.RateLimiter(1, 1, func(w http.ResponseWriter, r *http.Request) {
					// Write code in here
					res, _ := json.Marshal(model.ResponseMessage{Status: "success", Message: "message"})
					logging.Log("message", "INFO", r)
					w.WriteHeader(200)
					w.Write(res)
				}),
			)
		}
		`, module, module, module, capitalize(name))

	// this service template
	servTemp := fmt.Sprintf(`
		package service

		import (
			"database/sql"
		)

		type %sService struct {
			db *sql.DB
		}

		type %sService interface {
			ExampleService(id string) (string, error)
			// Add function in here
		}

		func New%sService(db *sql.DB) %sService {
			return &%sService{db}
		}

		// Write code in here
		func (q *%sService) ExampleService(id string) (string, error) {
			return "success", nil
		}
	`, name, capitalize(name), capitalize(name), capitalize(name), name, name)

	// this repository template

	repoTemp := fmt.Sprintf(`
		package repository

		import (
			"database/sql"
		)

		type %sRepository struct {
			db *sql.DB
		}

		type %sRepository interface {
			ExampleRepo(id string) error
			// Add function in here
		}

		func New%sRepository(db *sql.DB) %sRepository {
			return &%sRepository{db}
		}

		// Write code in here
		func (q *%sRepository) ExampleRepo(id string) error {

			if _, err := q.db.Exec("INSERT INTO table VALUES($1)", id); err != nil {
				return err
			}

			return nil
		}
		`, name, capitalize(name), capitalize(name), capitalize(name), name, name)

	switch ty {
	case "-h":
		handleTemp := process(hdlTemp, "handler", dir, name)
		fmt.Println(handleTemp)
		return
	case "-s":
		serviceTemp := process(servTemp, "service", dir, name)
		fmt.Println(serviceTemp)
		return
	case "-r":
		repositoryTemp := process(repoTemp, "repository", dir, name)
		fmt.Println(repositoryTemp)
		return
	case "-a":
		handleTemp := process(hdlTemp, "handler", dir, name)
		serviceTemp := process(servTemp, "service", dir, name)
		repositoryTemp := process(repoTemp, "repository", dir, name)
		fmt.Println(handleTemp)
		fmt.Println(serviceTemp)
		fmt.Println(repositoryTemp)
		return
	default:
		fmt.Println("invalid command")
		return
	}

}

func process(template, path, dir, name string) string {
	file := name + ".go"

	os.MkdirAll("internal/"+path+dir, 0o755)
	handlePath := "internal/" + path + dir
	save := filepath.Join(handlePath, file)

	os.WriteFile(save, []byte(template), 0o644)
	return "Created:" + save
}

func moduleReader() string {
	file, err := os.Open("go.mod")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "module ") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(line, "module"))
			return moduleName
		}
	}

	return ""
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}

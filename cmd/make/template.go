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

	// This handler template
	hdlTemp := fmt.Sprintf(`
package handler

import (
	"%s/pkg/middleware"
	"%s/pkg/response"
	"database/sql"
	"net/http"
)

// this handler is for URLs that do not have a prefix parameter
// example without prefix -> api/%s
func %sHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			handler(func(w http.ResponseWriter, r *http.Request) {
				// Write code in here
			})
		case http.MethodPost:
			handler(func(w http.ResponseWriter, r *http.Request) {
				// Write code in here
			})
		default:
			response.ResponseMessage("method not allowed", "method not allowed", "INFO", 405, w, r)
		}

	}
}

// example with prefix -> api/%s/{id}
func %sHandlerPrefix(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			handler(func(w http.ResponseWriter, r *http.Request) {
				// Write code in here
			})
		case http.MethodPut:
			handler(func(w http.ResponseWriter, r *http.Request) {
				// Write code in here
			})
		case http.MethodDelete:
			handler(func(w http.ResponseWriter, r *http.Request) {
				// Write code in here
			})
		default:
			response.ResponseMessage("method not allowed", "method not allowed", "INFO", 405, w, r)
		}

	}
}

// Custom your handler
func handler(next http.HandlerFunc) http.HandlerFunc {
	return middleware.CORS(
		middleware.RateLimiter(1, 1, func(w http.ResponseWriter, r *http.Request) {
			// add more middleware in here
			next(w, r)
		}),
	)
}
`, module, module, name, capitalize(name), name, capitalize(name))

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
	// Add function in here
	ExampleService(id string) (string, error)
}

func New%sService(db *sql.DB) %sService {
	return &%sService{db}
}

// Write code in here
func (q *%sService) ExampleService(id string) (string, error) {
	return "success", nil
}
`, name, capitalize(name), capitalize(name), capitalize(name), name, name)

	modeltemp := fmt.Sprintf(
		"package model\n\n"+
			"import %s\n\n"+
			"type %sResponse struct {\n"+
			"\tID   uuid.UUID `json:\"id\"`\n"+
			"\tName string `json:\"name\"`\n"+
			"}\n\n"+
			"type %sRequest struct {\n"+
			"\tName string `json:\"name\"`\n"+
			"}\n",
		`"github.com/google/uuid"`, capitalize(name), capitalize(name),
	)

	// this repository template
	repoTemp := fmt.Sprintf(`
package repository

import (
	"%s/internal/model"
	"database/sql"

	"github.com/google/uuid"
)

type %sRepository struct {
	db *sql.DB
}

type %sRepository interface {
	Show%s(page, size int, keyword string) ([]model.%sResponse, int, error)
	ShowById%s(id uuid.UUID) (model.%sResponse, error)
	Create%s(data model.%sRequest) error
	Update%s(data model.%sRequest, id string) error
	Delete%s(id uuid.UUID) error
	// Add function in here
}

func New%sRepository(db *sql.DB) %sRepository {
	return &%sRepository{db}
}

// Write code in here
func (q *%sRepository) Show%s(page, size int, keyword string) ([]model.%sResponse, int, error) {
	var count int
	if err := q.db.QueryRow("SELECT COUNT(*) FROM %s WHERE name ILIKE $1", %s+keyword+%s).Scan(&count); err != nil {
		return nil, 0, err
	}

	res, err := q.db.Query("SELECT id, name FROM %s WHERE name ILIKE $1 LIMIT $2 OFFSET $3", %s+keyword+%s, size, page)
	if err != nil {
		return nil, 0, err
	}

	defer res.Close()

	var result []model.%sResponse
	for res.Next() {
		var data model.%sResponse

		if err := res.Scan(&data.ID, &data.Name); err != nil {
			return nil, 0, err
		}

		result = append(result, data)
	}

	return result, count, nil
}

func (q *%sRepository) ShowById%s(id uuid.UUID) (model.%sResponse, error) {
	var data model.%sResponse
	if err := q.db.QueryRow("SELECT id name FROM %s WHERE id = $1", id).Scan(&data.ID, &data.Name); err != nil {
		return model.%sResponse{}, err
	}

	return data, nil
}

func (q *%sRepository) Create%s(data model.%sRequest) error {

	if _, err := q.db.Exec("INSERT INTO %s(name) VALUES($1)", data.Name); err != nil {
		return err
	}

	return nil
}

func (q *%sRepository) Update%s(data model.%sRequest, id string) error {

	if _, err := q.db.Exec("UPDATE %s SET name = $1 WHERE id = $2", data.Name, id); err != nil {
		return err
	}

	return nil
}

func (q *%sRepository) Delete%s(id uuid.UUID) error {

	if _, err := q.db.Exec("DELETE FROM %s WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}

`, module, name, capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), capitalize(name), name, name, capitalize(name), capitalize(name), name, `"%"`, `"%"`, name, `"%"`, `"%"`, capitalize(name), capitalize(name), name, capitalize(name), capitalize(name), capitalize(name), name, capitalize(name), name, capitalize(name), capitalize(name), name, name, capitalize(name), capitalize(name), name, name, capitalize(name), name)

	switch ty {
	case "-h", "handler":
		handleTemp := process(hdlTemp, "handler", dir, name)
		fmt.Println(handleTemp)
		return
	case "-s", "service":
		serviceTemp := process(servTemp, "service", dir, name)
		fmt.Println(serviceTemp)
		return
	case "-m", "model":
		modelTemp := process(modeltemp, "model", dir, name)
		fmt.Println(modelTemp)
		return
	case "-r", "repository":
		repositoryTemp := process(repoTemp, "repository", dir, name)
		fmt.Println(repositoryTemp)
		return
	case "-a", "all":
		handleTemp := process(hdlTemp, "handler", dir, name)
		serviceTemp := process(servTemp, "service", dir, name)
		modelTemp := process(modeltemp, "model", dir, name)
		repositoryTemp := process(repoTemp, "repository", dir, name)
		fmt.Println(handleTemp)
		fmt.Println(serviceTemp)
		fmt.Println(modelTemp)
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

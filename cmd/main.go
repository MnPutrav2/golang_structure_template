package main

import (
	"clean-arsitektur/internal/config"
	"clean-arsitektur/internal/route"
	"fmt"
	"net/http"
	"os"
)

func main() {
	db := config.Database()
	defer db.Close()

	// Route
	http.HandleFunc("/book", route.BookRoute(db))

	fmt.Println("[  Database connected ]")
	fmt.Println("[  Server listen in port ", os.Getenv("APP_LISTEN"), " ]")
	http.ListenAndServe(os.Getenv("APP_LISTEN"), nil)
}

package cmd

import (
	"clean-arsitektur/config"
	"fmt"
	"net/http"
	"os"
)

func Server(listen string) {
	db, err := config.Database()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Route
	// http.HandleFunc("/book", route.BookRoute(db))

	fmt.Println("[  Database connected ]")
	fmt.Println("[  Server listen in port ", os.Getenv("APP_LISTEN"), " ]")
	http.ListenAndServe(listen, nil)
}

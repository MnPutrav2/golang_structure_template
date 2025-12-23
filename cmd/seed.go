package cmd

import (
	"clean-arsitektur/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Seed() {
	fmt.Println("Running seed...")

	db, err := config.Database()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	entries, err := os.ReadDir("database/seed")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".sql") {
			continue
		}

		path := filepath.Join("database/seed", name)
		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("read %s: %v", name, err)
		}

		if _, err := db.Exec(string(sqlBytes)); err != nil {
			log.Fatalf("exec %s: %v", name, err)
		}

		fmt.Println("Seeded:", name)
	}

	fmt.Println("Seeding completed")
}

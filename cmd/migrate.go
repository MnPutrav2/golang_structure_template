package cmd

import (
	"clean-arsitektur/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Migrate() {
	fmt.Println("Running migrations... ==================")

	db, err := config.Database()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	entries, err := os.ReadDir("db/migrations")
	if err != nil {
		log.Fatal(err)
	}

	var ups []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if strings.HasSuffix(name, ".up.sql") {
			ups = append(ups, name)

		}
	}
	sort.Strings(ups)

	for _, name := range ups {
		path := filepath.Join("db/migrations", name)

		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("read %s: %v", name, err)
		}

		if _, err = db.Exec(string(sqlBytes)); err != nil {
			log.Fatalf("Error in %s: %v", name, err)
		}

		fmt.Println("Migrated:", name)
	}

	fmt.Println("Migration completed. ==================")
}

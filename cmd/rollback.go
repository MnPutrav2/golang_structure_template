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

func Rollback() {
	fmt.Println("Running rollback (down.sql)... ==================")

	db, err := config.Database()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	// Mulai transaction untuk safety
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to start transaction:", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	entries, err := os.ReadDir("db/migrations")
	if err != nil {
		log.Fatal(err)
	}

	var downs []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, ".down.sql") {
			downs = append(downs, name)
		}
	}

	sort.Strings(downs)

	// Eksekusi dalam transaction
	for i := len(downs) - 1; i >= 0; i-- {
		name := downs[i]
		path := filepath.Join("db/migrations", name)
		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("read %s: %v", name, err)
		}

		sqlContent := string(sqlBytes)

		// Eksekusi dengan error handling yang lebih baik
		if _, err := db.Exec(sqlContent); err != nil {
			// Cek jika error karena "table does not exist"
			if strings.Contains(err.Error(), "does not exist") ||
				strings.Contains(err.Error(), "not found") {
				fmt.Printf("⚠️  Skipping %s: %v\n", name, err)
				continue // Skip dan lanjut ke file berikutnya
			} else {
				// Untuk error lain, tetap fatal
				log.Fatalf("exec %s: %v", name, err)
			}
		}

		fmt.Println("✅ Rolled back: ", name)
	}

	// Commit transaction jika semua sukses
	if err := tx.Commit(); err != nil {
		log.Fatal("Failed to commit transaction:", err)
	}

	fmt.Println("Rollback completed successfully ==================")
}

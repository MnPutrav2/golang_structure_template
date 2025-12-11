package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	migrationsPath := "./db/migrations"
	db := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, addr, name)

	if len(os.Args) < 2 {
		fmt.Println("Usage example:")
		fmt.Println("  go run db/migrate.go create table_name")
		fmt.Println("  go run db/migrate.go up")
		fmt.Println("  go run db/migrate.go rollback")
		fmt.Println("  go run db/migrate.go steps -1")
		fmt.Println("  go run db/migrate.go force 2")
		return
	}

	cmd := os.Args[1]

	m, err := migrate.New("file://"+migrationsPath, db)
	if err != nil {
		log.Fatal("Error creating migrate instance:", err)
	}

	switch cmd {
	case "create":
		timestamp := time.Now().Format("20060102150405")

		upFile := filepath.Join(migrationsPath, fmt.Sprintf("%s_%s.up.sql", timestamp, os.Args[2]))
		downFile := filepath.Join(migrationsPath, fmt.Sprintf("%s_%s.down.sql", timestamp, os.Args[2]))

		upContent := "-- Write your UP migration here\n"
		downContent := "-- Write your DOWN migration here\n"

		if err := os.WriteFile(upFile, []byte(upContent), 0644); err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := os.WriteFile(downFile, []byte(downContent), 0644); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Created:")
		fmt.Println("  ", upFile)
		fmt.Println("  ", downFile)
	case "up":
		fmt.Println("Running UP...")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration error:", err)
		}
		fmt.Println("Migration UP completed.")

	case "rollback":
		fmt.Println("Running DOWN (1 step)...")
		if err := m.Steps(-1); err != nil {
			log.Fatal("Rollback error:", err)
		}
		fmt.Println("Rollback DOWN completed.")

	case "steps":
		if len(os.Args) < 3 {
			log.Fatal("Missing steps argument. Example: steps -1")
		}
		steps, _ := strconv.Atoi(os.Args[2])
		fmt.Println("Running STEPS:", steps)
		if err := m.Steps(steps); err != nil {
			log.Fatal("Steps error:", err)
		}
		fmt.Println("Steps migration completed.")

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Missing version. Example: force 2")
		}
		version, _ := strconv.Atoi(os.Args[2])
		fmt.Println("Forcing version:", version)
		if err := m.Force(version); err != nil {
			log.Fatal("Force error:", err)
		}
		fmt.Println("Force version completed.")

	default:
		fmt.Println("Unknown command:", cmd)
	}
}

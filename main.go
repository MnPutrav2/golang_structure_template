package main

import (
	"clean-arsitektur/cmd"
	"clean-arsitektur/cmd/make"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 2 {
		Help()
		return
	}

	cli := os.Args[1]
	args := os.Args[2:]

	prod := os.Getenv("APP_LISTEN")
	dev := os.Getenv("APP_LISTEN_DEV")

	switch cli {
	case "server":
		cmd.Server(prod)

	case "dev":
		cmd.Server(dev)

	case "migrate":
		cmd.Migrate()

	case "rollback":
		cmd.Rollback()

	case "seed":
		cmd.Seed()

	case "fresh":
		cmd.Fresh()

	case "make:migration":
		if len(args) == 0 {
			fmt.Println("Usage: go run . make:migration <Name>")
			return
		}
		make.Migration(args[0])

	case "scheduler":
		cmd.Schduler()

	case "make:seeder":
		if len(args) == 0 {
			fmt.Println("Usage: go run . make:seeder <Name>")
			return
		}
		make.Seeder(args[0])

	case "make:template":
		if len(args) == 0 || len(args[1]) == 0 || len(args[2]) == 0 {
			fmt.Println("Usage: go run . make:template <folder_name> <file_name> <type>")
			return
		}

		make.Template(args[0], args[1], args[2])

	default:
		fmt.Print("Command not found")
	}
}

func Help() {
	fmt.Print(`
Available commands:

	Application:
		go run . server							run server

	Database migration:
		go run . migrate						run migration
		go run . rollback						down migration
		go run . seed							run seed SQL
		go run . fresh							rollback -> migrate -> seed
		go run . make:migration <table_name>				create timestamped up/down migration pair
		go run . make:seed <table_name>					create seed SQL template
		
	Create template:
		go run . make:template <create_folder> <file_name> <type>	create template, <type> = [-h = create handler, -ro = create route, -r = create repository, -s = create service, -a = create all]

	Scheduler:
		go run . scheduler						scheduler job, edit cmd/scheduler.go
	`)
}

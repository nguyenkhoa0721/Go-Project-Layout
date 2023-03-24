package main

import (
	"bytes"
	"fmt"
	"github.com/nguyenkhoa0721/go-project-layout/config"
	"os"
	"os/exec"
)

const ShellToUse = "bash"

func Shellout(command string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	fmt.Println(command)
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
	return err
}

func main() {
	config, _ := config.LoadConfig()

	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument.")
		os.Exit(1)
	}

	arg := os.Args[1]

	if arg == "migrateup" {
		err := Shellout(fmt.Sprintf("migrate -path pkg/db/postgres/migration -database \"postgres://%s:%s@%s:%d/%s?sslmode=disable\" -verbose up", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Database))
		if err != nil {
			panic(err)
		}
	} else if arg == "migratedown" {
		err := Shellout(fmt.Sprintf("migrate -path pkg/db/postgres/migration -database \"postgres://%s:%s@%s:%d/%s?sslmode=disable\" -verbose down", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Database))
		if err != nil {
			panic(err)
		}
	} else if arg == "sqlc" {
		err := Shellout("sqlc generate")
		if err != nil {
			panic(err)
		}
	} else if arg == "dev" {
		err := Shellout("docker compose -p explorer -f build/dev/docker-compose.yml down")
		err = Shellout(fmt.Sprintf("export BITBUCKET_TOKEN=%s && docker compose -p explorer -f build/dev/docker-compose.yml build --no-cache", config.Build.GitToken))
		err = Shellout("docker compose -p explorer -f build/dev/docker-compose.yml up -d")
		if err != nil {
			panic(err)
		}
	}
}

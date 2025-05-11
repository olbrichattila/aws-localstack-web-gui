// Package main is the main entrypoint
package main

import (
	"api/internal/database"
	"api/internal/server"
	"fmt"
	"os"
)

func main() {
	db, err := database.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server, err := server.New(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = server.Serve(8080)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

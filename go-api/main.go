// Package main is the main entrypoint
package main

import (
	"api/internal/aws/awsshared"
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

	awsShared := awsshared.New(db)

	server, err := server.New(db, awsShared)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Starting server on port 8080")
	err = server.Serve(8080)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Package main is the main entrypoint
package main

import (
	"api/internal/aws/awsshared"
	"api/internal/database"
	"api/internal/server"
	"fmt"
	"os"
)

const runPort = 80

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

	fmt.Printf("Starting server on port %d\n", runPort)
	err = server.Serve(runPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

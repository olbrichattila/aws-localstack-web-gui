package server

import (
	"api/internal/aws/awsshared"
	"api/internal/database"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	errServer = errors.New("HTTP Server error")
)

// Server encapsulates logic to serve on HTTP
type Server interface {
	Serve(port int) error
}

type server struct {
	db        database.Database
	awsShared awsshared.AWSShared
}

// New creates a new HTTP server
func New(db database.Database, awsShared awsshared.AWSShared) (Server, error) {
	if db == nil {
		return nil, fmt.Errorf("Database is nil %w", errServer)
	}

	if awsShared == nil {
		return nil, fmt.Errorf("awsShared is nil %w", errServer)
	}

	return &server{
		db:        db,
		awsShared: awsShared,
	}, nil
}

// Serve serves on HTTP with the provided port
func (s *server) Serve(port int) error {
	s.initRoutes()

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		return fmt.Errorf("Could not spin up the server %w", errServer)
	}

	return nil
}

package snslistener

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"webuiApi/app/repositories/database"
)

// SNSRequest represents an SNS subscription or notification request

type SNSRequest struct {
	Type              string          `json:"Type"`
	SubscribeURL      string          `json:"SubscribeURL,omitempty"`
	Token             string          `json:"Token,omitempty"`
	TopicArn          string          `json:"TopicArn,omitempty"`
	Message           string          `json:"Message,omitempty"`
	MessageID         string          `json:"MessageId,omitempty"`
	Timestamp         string          `json:"Timestamp,omitempty"`
	SignatureVersion  string          `json:"SignatureVersion,omitempty"`
	Signature         string          `json:"Signature,omitempty"`
	SigningCertURL    string          `json:"SigningCertURL,omitempty"`
	MessageAttributes json.RawMessage `json:"MessageAttributes,omitempty"`
}

func New() SNSListener {
	return &snsListener{
		listeners:      make(map[int]*listener),
		requestHistory: make(map[int][]SNSRequest),
		mutex:          &sync.RWMutex{},
	}
}

type ListenerInfo struct {
	Port int    `json:"port"`
	Info string `json:"info"`
}

type SNSListener interface {
	Construct(db database.Database)
	Listen(port int) error
	Close(port int) error
	Purge(port int) error
	GetRequests(port int) ([]SNSRequest, error)
	GetListeningPorts() []ListenerInfo
}

type listener struct {
	server *http.Server
	info   string
}

type snsListener struct {
	db             database.Database
	listeners      map[int]*listener
	requestHistory map[int][]SNSRequest
	mutex          *sync.RWMutex
}

// Construct implements SNSListener.
func (s *snsListener) Construct(db database.Database) {
	s.db = db
}

// Listen implements SNSListener.
func (s *snsListener) Listen(port int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.listeners[port]; ok {
		return fmt.Errorf("listener already open on port %d", port)
	}

	// Initialize request history for this port
	s.requestHistory[port] = []SNSRequest{}

	// Create a new ServeMux for this server
	mux := http.NewServeMux()

	// Handle SNS subscription and notification requests
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Parse SNS request
		var snsReq SNSRequest
		if err := json.Unmarshal(body, &snsReq); err != nil {
			http.Error(w, "Invalid SNS request format", http.StatusBadRequest)
			return
		}

		// Store the request
		s.mutex.Lock()
		s.requestHistory[port] = append(s.requestHistory[port], snsReq)
		s.mutex.Unlock()

		// Handle subscription confirmation
		if snsReq.Type == "SubscriptionConfirmation" && snsReq.Token != "" {
			// Return 200 OK with the token to confirm subscription
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(snsReq.Token))

			// Call subscribe URL
			setting, err := s.db.GetSettings()
			if err != nil {
				s.mutex.Lock()
				if _, ok := s.listeners[port]; ok {
					s.listeners[port].info = err.Error()
				}
				s.mutex.Unlock()
			}

			alteredSubURL := strings.Replace(snsReq.SubscribeURL, "http://localhost.localstack.cloud:4566", setting.Endpoint, -1)

			_, err = http.Get(alteredSubURL)
			if err != nil {
				s.mutex.Lock()
				if _, ok := s.listeners[port]; ok {
					s.listeners[port].info = err.Error()
				}
				s.mutex.Unlock()
			}

			return
		}

		// For other requests, just return 200 OK
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}

	s.listeners[port] = &listener{server: server, info: "Added"}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.mutex.Lock()
			if _, ok := s.listeners[port]; ok {
				s.listeners[port].info = err.Error()
			}
			s.mutex.Unlock()
		}
	}()

	return nil
}

// Close implements SNSListener.
func (s *snsListener) Close(port int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	listener, ok := s.listeners[port]
	if !ok {
		return fmt.Errorf("no listener on port %d", port)
	}

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := listener.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	// Remove from listeners map
	delete(s.listeners, port)
	return nil
}

// Close implements SNSListener.
func (s *snsListener) Purge(port int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.requestHistory, port)

	return nil
}

// GetRequests returns the request history for a specific port
func (s *snsListener) GetRequests(port int) ([]SNSRequest, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	requests, ok := s.requestHistory[port]
	if !ok {
		return nil, fmt.Errorf("no request history for port %d", port)
	}

	return requests, nil
}

// GetListeningPorts returns a sorted slice of all ports currently being listened to
func (s *snsListener) GetListeningPorts() []ListenerInfo {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	ports := make([]ListenerInfo, 0, len(s.listeners))
	for port := range s.listeners {
		ports = append(ports, ListenerInfo{Port: port, Info: s.listeners[port].info})
	}

	// Sort ports in ascending order
	for i := range len(ports) - 1 {
		for j := range len(ports) - i - 1 {
			if ports[j].Port > ports[j+1].Port {
				ports[j], ports[j+1] = ports[j+1], ports[j]
			}
		}
	}

	return ports
}

package middleware

import (
	"net/http"
)

// CorsMiddleware function can take any parameters defined in the Di config
func CorsMiddleware(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// time.Sleep(time.Second * 2) // TODO, test remove
}

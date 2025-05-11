package server

import "net/http"

type handleFuncWithReturn func(w http.ResponseWriter, r *http.Request) bool

func (*server) initMiddlewares(handleFunc http.HandlerFunc, middlewares ...handleFuncWithReturn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewares {
			ok := middleware(w, r)
			if !ok {
				return
			}
		}
		handleFunc(w, r)
	}
}

func (*server) corsMiddleware(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodOptions {
		return false
	}

	return true
}

func (*server) jSONMiddleware(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Content-Type", "application/json")

	return true
}

func (*server) getMiddleware(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func (*server) postMIddleware(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func (*server) getPostMiddleware(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodPost || r.Method == http.MethodGet {
		return true
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return false
}

func (*server) getPostDeleteMiddleware(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodPost || r.Method == http.MethodGet || r.Method == http.MethodDelete {
		return true
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return false
}

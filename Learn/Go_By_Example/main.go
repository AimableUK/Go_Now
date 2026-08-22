package main

import (
	"net/http"
)

type server struct {
	addr string
}

// V1
// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {

// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("Index Page"))
// 			return
// 		case "/users":
// 			w.Write([]byte("Users Page"))
// 			return
// 		}

// 	default:
// 		w.Write([]byte("404 Page"))
// 		return
// 	}
// }

// V2
type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("GET method"))
		case "/index":
			w.Write([]byte("GET index"))
		}
	case http.MethodPost:
		w.Write([]byte("POST method"))
	}
}

func main() {

	// V1
	// s := &server{addr: ":8080"}

	// if err := http.ListenAndServe(s.addr, s); err != nil {
	// 	log.Fatal(err)
	// }

	// V2
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("/users", api.getUsersHandler)

	srv.ListenAndServe()
}

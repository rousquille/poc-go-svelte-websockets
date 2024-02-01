package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rousquille/poc-go-svelte-websockets/internal/cli"
	"log"
	"net/http"
)

type Server struct {
	router    *mux.Router
	wsManager *WsManager
}

func newServer() *Server {
	ws := NewWsManager()

	s := &Server{
		router:    mux.NewRouter(),
		wsManager: ws,
	}

	s.routes()

	if *cli.Cors {
		s.router.Use(corsMiddleware)
	}

	return s
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	logRequestMiddleware(s.router.ServeHTTP).ServeHTTP(w, r)
}

func RunWebServer(url string) error {
	var err error
	srv := newServer()

	http.HandleFunc("/", srv.serveHTTP)
	log.Println("Serving HTTP on", url)

	err = http.ListenAndServe(url, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) decodeJSON(_ http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) respondJSON(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

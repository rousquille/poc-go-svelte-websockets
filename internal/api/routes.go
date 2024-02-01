package api

import "net/http"

func (s *Server) routes() {
	s.router.HandleFunc("/ws", s.wsHandler()).Methods("GET")
	s.router.PathPrefix("/").Handler(http.FileServer(http.FS(FrontendContents())))
}

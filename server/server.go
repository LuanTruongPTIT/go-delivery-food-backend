package server

import "net/http"

type Server struct {
	srv *http.Server
}

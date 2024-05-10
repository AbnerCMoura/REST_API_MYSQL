package api

import (
	"database/sql"
	"fmt"
	"github.com/AbnerCMoura/REST_API_MYSQL/services/user"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(adrr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: adrr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)
	
	fmt.Println("Listening on " + s.addr)

	return http.ListenAndServe(s.addr, router)
}

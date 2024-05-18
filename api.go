package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	Addr string
}

type apifunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func writeJson(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

func httpHandleFunc(f apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		Addr: addr,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", httpHandleFunc(s.handleAccount)).Methods("GET", "POST", "DELETE")

	log.Println("Server is running on port :", s.Addr)
	http.ListenAndServe(s.Addr, router)
}

func (s *ApiServer) handleAccount(wr http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(wr, r)
	case "POST":
		return s.handleCreateAccount(wr, r)
	case "DELETE":
		return s.handleDeleteAccount(wr, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *ApiServer) handleGetAccount(wr http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleCreateAccount(wr http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(wr http.ResponseWriter, r *http.Request) error {
	return nil
}


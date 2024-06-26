package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	Addr string
	Db   Db
}

type apifunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
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

func NewApiServer(addr string, db Db) *ApiServer {
	return &ApiServer{
		Addr: addr,
		Db:   db,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", httpHandleFunc(s.handleAccount)).Methods("GET", "POST", "DELETE")

	router.HandleFunc("/account/{id}", httpHandleFunc(s.handleGetAccountByID)).Methods("GET")

	router.HandleFunc("/account/{id}", httpHandleFunc(s.handleDeleteAccount)).Methods("DELETE")

	log.Println("Server is running on port :", s.Addr)
	http.ListenAndServe(s.Addr, router)
}

func (s *ApiServer) handleAccount(wr http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccounts(wr, r)
	case "POST":
		return s.handleCreateAccount(wr, r)
	case "DELETE":
		return s.handleDeleteAccount(wr, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

func (s *ApiServer) handleGetAccountByID(wr http.ResponseWriter, r *http.Request) error {

	id, err := getId(r)
	if err != nil {
		return err
	}

	acc, err := s.Db.GetAccountByID(id)

	if err != nil {
		return err
	}
	return writeJson(wr, http.StatusOK, acc)

}

func (s *ApiServer) handleGetAccounts(wr http.ResponseWriter, r *http.Request) error {
	account, err := s.Db.GetAccounts()

	if err != nil {
		return err
	}

	return writeJson(wr, http.StatusOK, account)

}

func (s *ApiServer) handleCreateAccount(wr http.ResponseWriter, r *http.Request) error {
	createAccReq := new(CreateAccountReq)

	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
		return err
	}
	account := NewAccount(createAccReq.FirstName, createAccReq.LastName)

	if err := s.Db.CreateAccount(account); err != nil {
		return err
	}

	return writeJson(wr, http.StatusOK, account)
}

func (s *ApiServer) handleDeleteAccount(wr http.ResponseWriter, r *http.Request) error {

	id, err := getId(r)
	if err != nil {
		return err
	}

	err = s.Db.DeleteAccount(id)

	if err != nil {
		return err
	}
	return writeJson(wr, http.StatusOK, "Deletion successful")
}

func getId(r *http.Request) (int, error) {
	idstr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idstr)

	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idstr)
	}

	return id, nil

}

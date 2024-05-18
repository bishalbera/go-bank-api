package main

import "net/http"

type ApiServer struct {
	Addr string
}

func newApiServer(addr string) *ApiServer {

	return &ApiServer{
		Addr: addr,
	}
}

func (s *ApiServer) handleAccount(wr http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleGetAccount(wr http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleCreateAccount(wr http.ResponseWriter, r *http.Request) error {
	return nil
}
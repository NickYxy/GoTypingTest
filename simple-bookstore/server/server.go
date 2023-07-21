package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-bookstore/server/middleware"
	"simple-bookstore/store"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := &BookStoreServer{
		s: s,
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.updateBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.getBookHandler).Methods("GET")
	router.HandleFunc("/book", srv.getAllBooksHandler).Methods("GET")
	router.HandleFunc("/book/{id}", srv.deleteBookHandler).Methods("DELETE")

	srv.srv.Handler = middleware.Logging(middleware.Validating(router))
	return srv
}

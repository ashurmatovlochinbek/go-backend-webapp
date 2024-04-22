package router

import (
	"simple-go-app/internal/handler"

	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, sh *handler.StudentHandler) {
	r.HandleFunc("/students", sh.GetAllStudents).Methods("GET")
	r.HandleFunc("/student/{id}", sh.GetById).Methods("GET")
	r.HandleFunc("/student", sh.Create).Methods("POST")
	r.HandleFunc("/student/{id}", sh.Update).Methods("PUT")
	r.HandleFunc("/student/{id}", sh.Delete).Methods("DELETE")
}

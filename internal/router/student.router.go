package router

import (
	"simple-go-app/internal/handler"

	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router, sh *handler.StudentHandler) {
	r.HandleFunc("/students", sh.GetAllStudents).Methods("GET")
}

package handler

import (
	"encoding/json"
	"net/http"
	"simple-go-app/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	StudentService service.StudentService
}

func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.StudentService.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonFormat, err := json.Marshal(students)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFormat)
}

func (h *StudentHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student, err := h.StudentService.GetById(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonFormat, err := json.Marshal(student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFormat)
}

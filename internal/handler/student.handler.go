package handler

import (
	"encoding/json"
	"net/http"
	"simple-go-app/internal/service"
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

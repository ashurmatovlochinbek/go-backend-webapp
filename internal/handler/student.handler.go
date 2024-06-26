package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"simple-go-app/internal/model"
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
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.StudentService.Create(r.Context(), &student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(id)))
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var student model.Student
	err = json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedStudent, err := h.StudentService.Update(r.Context(), &student, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonFormat, err := json.Marshal(updatedStudent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonFormat)
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	affectedRows, err := h.StudentService.Delete(r.Context(), id)

	if affectedRows == 0 {
		http.Error(w, "user not found", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(affectedRows)))
}

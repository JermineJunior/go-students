package handlers

import (
	"net/http"

	"github.com/JermaineJunior/goStudents/internal/services"
	"github.com/JermaineJunior/goStudents/views/pages"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{
		service: service,
	}
}

// handlers
func (h *StudentHandler) Index(w http.ResponseWriter, r *http.Request) error {
	students, err := h.service.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return Render(w, r, pages.Students(students))
}

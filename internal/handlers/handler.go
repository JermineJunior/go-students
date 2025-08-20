package handlers

import (
	"net/http"

	"github.com/JermaineJunior/goStudents/views/pages"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	err := Render(w, r, pages.Home())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}
func HandleNotFound(w http.ResponseWriter,r *http.Request) error {
	return Render(w,r,pages.NotFound())
}
package services

import (
	"github.com/JermaineJunior/goStudents/internal/models"
	"github.com/JermaineJunior/goStudents/internal/records"
)

type StudentService struct {
	record *records.StudentRecord
}

func NewStudentService(record *records.StudentRecord) *StudentService {
	return &StudentService{
		record: record,
	}
}

func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	return s.record.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id int) (models.Student, error) {
	return s.record.GetStudentByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) (int64, error) {
	return s.record.CreateStudent(student)
}

func (s *StudentService) UpdateStudent(id int, student models.Student) error {
	return s.record.DeleteStudent(id)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.record.DeleteStudent(id)
}

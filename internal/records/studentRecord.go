package records

import (
	"database/sql"

	"github.com/JermaineJunior/goStudents/internal/models"
)

type StudentRecord struct {
	db *sql.DB
}

func NewStudentRecord(db *sql.DB) *StudentRecord {
	return &StudentRecord{
		db: db,
	}
}

func (r *StudentRecord) GetAllStudents() ([]models.Student, error) {
	rows, err := r.db.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []models.Student{}
	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Address, &student.CreatedAt)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *StudentRecord) GetStudentByID(id int) (models.Student, error) {
	row := r.db.QueryRow("SELECT * FROM students WHERE id = ?", id)
	var student models.Student
	err := row.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Address, &student.CreatedAt)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentRecord) CreateStudent(student models.Student) (int64, error) {
	result, err := r.db.Exec("INSERT INTO students (name, email, phone, address) VALUES (?, ?, ?, ?)", student.Name, student.Email, student.Phone, student.Address)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *StudentRecord) UpdateStudent(id int, student models.Student) error {
	_, err := r.db.Exec("UPDATE students SET name = ?, email = ?, phone = ?, address = ? WHERE id = ?", student.Name, student.Email, student.Phone, student.Address, id)
	return err
}

func (r *StudentRecord) DeleteStudent(id int) error {
	_, err := r.db.Exec("DELETE FROM students WHERE id = ?", id)
	return err
}

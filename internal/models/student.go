package models

import "time"

type Student struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Address   string
	CreatedAt time.Time
}

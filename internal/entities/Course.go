package entities

import "time"

type Course struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  string    `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateddAt  time.Time `json:"updatedAt"`
}

func NewCourse() *Course {
	return &Course{}
}

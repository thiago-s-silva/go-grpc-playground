package entities

import (
	"errors"
	"time"
)

type Course struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  string    `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewCourse() *Course {
	return &Course{}
}

func (c *Course) Validate() error {
	if c.ID == "" {
		return errors.New("the course id should be defined")
	}

	return nil
}

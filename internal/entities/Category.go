package entities

import "time"

type Category struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateddAt  time.Time `json:"updatedAt"`
}

func NewCategory() *Category {
	return &Category{}
}

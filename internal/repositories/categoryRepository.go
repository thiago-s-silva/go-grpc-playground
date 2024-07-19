package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thiago-s-silva/grpc-example/internal/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(name string, description string) (*entities.Category, error)
	FindAll() ([]*entities.Category, error)
	FindByCourseID(courseID string) (*entities.Category, error)
	FindByID(id string) (*entities.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(name string, description string) (*entities.Category, error) {
	category := entities.Category{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}

	if res := r.db.Create(&category); res.Error != nil {
		return nil, res.Error
	}

	return &category, nil
}

func (r *categoryRepository) FindAll() ([]*entities.Category, error) {
	var categories []*entities.Category

	if res := r.db.Find(&categories); res.Error != nil {
		return nil, res.Error
	}

	return categories, nil
}

func (r *categoryRepository) FindByCourseID(courseID string) (*entities.Category, error) {
	var category entities.Category

	res := r.db.Raw(
		"SELECT * FROM categories WHERE course_id = ?",
		courseID,
	).Scan(&category)
	if res.Error != nil {
		return nil, res.Error
	}

	return &category, nil
}

func (r *categoryRepository) FindByID(id string) (*entities.Category, error) {
	var category entities.Category

	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thiago-s-silva/grpc-example/internal/entities"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(name string, description string) (*entities.Course, error)
	FindAll() ([]*entities.Course, error)
	FindByCategoryID(categoryId string) (*entities.Course, error)
	FindByID(id string) (*entities.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *courseRepository {
	return &courseRepository{
		db: db,
	}
}

func (r *courseRepository) Create(name string, description string) (*entities.Category, error) {
	course := entities.Category{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}

	if res := r.db.Create(&course); res.Error != nil {
		return nil, res.Error
	}

	return &course, nil
}

func (r *courseRepository) FindAll() ([]*entities.Course, error) {
	var courses []*entities.Course

	if res := r.db.Find(&courses); res.Error != nil {
		return nil, res.Error
	}

	return courses, nil
}

func (r *courseRepository) FindByCategoryID(categoryId string) (*entities.Category, error) {
	var course entities.Category

	res := r.db.Raw(
		"SELECT * FROM courses WHERE category_id = ?",
		categoryId,
	).Scan(&course)
	if res.Error != nil {
		return nil, res.Error
	}

	return &course, nil
}

func (r *courseRepository) FindByID(id string) (*entities.Category, error) {
	var course entities.Category

	err := r.db.Where("id = ?", id).First(&course).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &course, nil
}

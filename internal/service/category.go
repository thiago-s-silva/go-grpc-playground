package service

import (
	"context"
	"fmt"
	"io"

	"github.com/thiago-s-silva/grpc-example/internal/pb"
	"github.com/thiago-s-silva/grpc-example/internal/repositories"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	fmt.Println("CreateCategory was invoke")

	// Create the category on db getting the data from the gRPC request payload
	category, err := c.categoryRepository.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	// Define the Response DTO that's a category defined on proto file
	res := pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &res, nil
}

func (c *CategoryService) ListCategories(context.Context, *pb.Blank) (*pb.CategoryList, error) {
	// Get all categories from DB
	dbCategories, err := c.categoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	// Define a new variable to store the pb categories (gRPC format)
	var categories []*pb.Category

	// For each db category found
	for _, category := range dbCategories {
		// Create a new variable to parse the db category to a pb category
		categoryRes := pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		// Append to the pb categories list
		categories = append(categories, &categoryRes)
	}

	// Define a res variable following the pb service DTO response
	res := &pb.CategoryList{
		Categories: categories,
	}

	return res, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryByIdRequest) (*pb.Category, error) {
	// Try to get the category from DB using the received ID
	categoryDb, err := c.categoryRepository.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	// Check if the category was not found
	if categoryDb == nil {
		// Return an error message
		return nil, fmt.Errorf("category not found")
	}

	// Define the response based on gRPC service DTO
	res := &pb.Category{
		Id:          categoryDb.ID,
		Name:        categoryDb.Name,
		Description: categoryDb.Description,
	}

	return res, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	// Define a list of categories using gRPC DTO format
	categories := &pb.CategoryList{}

	// Loop to listen the requests from the stream
	for {
		category, err := stream.Recv()
		// Validate if the stream has ended (there is no more data to be sent)
		if err == io.EOF {
			// Close the stream sending all mapped categories
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		// Create a new category on DB
		createdCategory, err := c.categoryRepository.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		// Append the created category on the stream list
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          createdCategory.ID,
			Name:        createdCategory.Name,
			Description: createdCategory.Description,
		})
	}
}

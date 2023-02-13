package service

import (
	"grpc_go/internal/database"
	"grpc_go/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.categoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	
	categoryResponse := &pb.Category{
		Id: 			category.Id,
		Name: 			category.Name,
		Description: 	category.Description
	}
	
	return &pb.CategoryResponse {
		Category: categoryResponse,
		}, nil
		
	}
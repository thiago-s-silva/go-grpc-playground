package main

import (
	"fmt"
	"net"

	"github.com/thiago-s-silva/grpc-example/internal/database"
	"github.com/thiago-s-silva/grpc-example/internal/pb"
	"github.com/thiago-s-silva/grpc-example/internal/repositories"
	"github.com/thiago-s-silva/grpc-example/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := database.InitializeSQLite()
	if err != nil {
		panic(fmt.Errorf("can't initialzie the sqlite db: %e", err))
	}

	// REPOSITORIES
	categoryRepository := repositories.NewCategoryRepository(db)
	// courseRepository := repositories.NewCourseRepository(db)

	// SERVICES
	categoryService := service.NewCategoryService(categoryRepository)

	// GRPC Server
	// init the gRPC server
	grpcServer := grpc.NewServer()

	// register the category service to the gRPC server
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// register the reflection service
	reflection.Register(grpcServer)

	// TCP Connection
	// open a new TCP connection for the default gRPC port
	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// try to connect gRPC to the openned TCP connection
	if err := grpcServer.Serve(list); err != nil {
		panic(err)
	}
}

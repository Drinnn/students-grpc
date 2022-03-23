package services

import (
	"context"

	"github.com/Drinnn/students-grpc/protos"
)

type StudentService struct {
	protos.UnimplementedStudentServiceServer
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (*StudentService) AddStudent(ctx context.Context, req *protos.Student) (*protos.Student, error) {
	return &protos.Student{
		Id: "123",
		Name: req.Name,
		Email: req.Email,
	}, nil
}
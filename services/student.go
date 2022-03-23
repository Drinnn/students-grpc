package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

func (*StudentService) AddStudentVerbose(req *protos.Student, stream protos.StudentService_AddStudentVerboseServer) error {
	stream.Send(&protos.StudentResultStream{
		Status: "Init",
		Student: &protos.Student{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&protos.StudentResultStream{
		Status: "Inserting",
		Student: &protos.Student{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&protos.StudentResultStream{
		Status: "User inserted",
		Student: &protos.Student{
			Id: "123",
			Name: req.Name,
			Email: req.Email,
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&protos.StudentResultStream{
		Status: "Completed",
		Student: &protos.Student{
			Id: "123",
			Name: req.Name,
			Email: req.Email,
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*StudentService) AddStudents(stream protos.StudentService_AddStudentsServer) error {
	students := []*protos.Student{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&protos.Students{
				Students: students,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		students = append(students, &protos.Student{
			Id: req.Id,
			Name: req.Name,
			Email: req.Email,
		})

		fmt.Println("Adding", req.Name)
	}
}
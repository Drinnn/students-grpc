package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Drinnn/students-grpc/protos"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := protos.NewStudentServiceClient(connection)
	
	// AddStudent(client)
	// AddStudentVerbose(client)
	AddStudents(client)
}

func AddStudent(client protos.StudentServiceClient) {
	req := &protos.Student{
		Name: "John",
		Email: "john@mail.com",
	}

	res, err := client.AddStudent(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddStudentVerbose(client protos.StudentServiceClient) {
	req := &protos.Student{
		Name: "John",
		Email: "john@mail.com",
	}

	responseStream, err := client.AddStudentVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not make gRPC request: %v", err)
		}
		
		fmt.Println("Status: ", stream.Status)
	}

}

func AddStudents(client protos.StudentServiceClient) {
	reqs := []*protos.Student{
		{
			Id: "1",
			Name: "John",
			Email: "john@mail.com",
		},
		{
			Id: "2",
			Name: "Adam",
			Email: "adam@mail.com",
		},
		{
			Id: "3",
			Name: "Doe",
			Email: "doe@mail.com",
		},
		{
			Id: "4",
			Name: "Abdhu",
			Email: "abdhu@mail.com",
		},
		{
			Id: "5",
			Name: "Wesley",
			Email: "wesley@mail.com",
		},
	}

	stream, err := client.AddStudents(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
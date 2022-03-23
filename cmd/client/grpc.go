package main

import (
	"context"
	"fmt"
	"io"
	"log"

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
	AddStudentVerbose(client)
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
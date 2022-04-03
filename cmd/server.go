package main

import (
	"encoding/json"
	"example/grpc/pb"
	"example/grpc/service"
	"example/model"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var courseList = model.NewCourses()

// web server returning the course list via REST
// inserting courses via gRPC
func main() {
	// starts listening in a different thred (go routing)
	go startGrpc()

	http.HandleFunc("/course", CourseListHandler)
	http.ListenAndServe(":8888", nil)
}

func startGrpc() {
	// create listener to open conection
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		panic(err)
	}

	// create gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// instantiate service object
	courseService := service.NewCourseGrpcService()
	courseService.Courses = courseList

	// register service on gRPC server
	pb.RegisterCourseServiceServer(grpcServer, courseService)
	grpcServer.Serve(listener)
}

func CourseListHandler(w http.ResponseWriter, r *http.Request) {
	courseJson, _ := json.Marshal(courseList)
	w.Write([]byte(courseJson))
}

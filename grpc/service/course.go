package service

import (
	"context"
	"example/grpc/pb"
	"example/model"
)

func NewCourseGrpcService() *CourseGrpcService {
	return &CourseGrpcService{}
}

type CourseGrpcService struct {
	pb.UnimplementedCourseServiceServer
	Courses *model.Courses
}

func (c *CourseGrpcService) CreateCourse(ctx context.Context, request *pb.Course) (*pb.CourseResult, error) {
	course := model.NewCourse()
	course.Name = request.Name
	c.Courses.Add(course)

	return &pb.CourseResult{
		Id:   course.ID,
		Name: course.Name,
	}, nil
}

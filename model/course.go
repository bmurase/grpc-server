package model

import uuid "github.com/satori/go.uuid"

type Course struct {
	ID   string
	Name string
}

type Courses struct {
	Course []Course
}

func (c *Courses) Add(course *Course) {
	c.Course = append(c.Course, *course)
}

func NewCourse() *Course {
	course := Course{}
	course.ID = uuid.NewV4().String()
	return &course
}

func NewCourses() *Courses {
	return &Courses{}
}

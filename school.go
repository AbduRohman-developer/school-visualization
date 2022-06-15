package main

// version 0.2

type school struct {
	name     string
	director director
	cls      []class
}

type director struct {
	name string
	age  string
}

type teacher struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	ClassName string `json:"className"`
}

type class struct {
	Name        string `json:"name"`
	TeacherName string `json:"teacherName"`
}

type pupil struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Mark        int    `json:"mark"`
	ParentsName string `json:"parentsName"`
	ClassName   string `json:"className"`
}

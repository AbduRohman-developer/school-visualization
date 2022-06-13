package main

type school struct {
	name     string
	director director
	classes  []class
}

type director struct {
	name string
	age  string
}

type teacher struct {
	name      string
	age       string
	className string
}

type class struct {
	name        string
	teacherName string
}

type classes []class

type pupil struct {
	name        string
	age         string
	mark        string
	parentsName string
	className   string
}

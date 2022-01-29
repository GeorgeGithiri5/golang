package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	speciality string
}

func main() {
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
}

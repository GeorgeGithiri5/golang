package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)

	fmt.Println("Mark is not an athlete any more")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}

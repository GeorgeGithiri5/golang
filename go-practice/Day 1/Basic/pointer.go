package main

import "fmt"

func add_by_value(a int) int {
	a = a + 1
	return a
}

func add_by_reference(a *int) int {
	*a = *a + 1
	return *a
}

func show_add_by_value() {
	x := 3
	fmt.Println("x = ", x)
	fmt.Println("add_by_value(x)", add_by_value(x))
	fmt.Println("x = ", x)
}

func show_add_by_reference() {
	x := 3
	fmt.Println("x = ", x)

	fmt.Println("add_by_reference(&x)= ", add_by_reference(&x))
	fmt.Println("x = ", x)
}

func main() {
	show_add_by_value()
	show_add_by_reference()
}

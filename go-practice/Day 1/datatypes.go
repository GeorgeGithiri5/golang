// manipulation of basic data types

package main

import (
	"fmt"
)

const Pi = 3.14115926

var isActive bool
var enabled, disabled = true, false

// grouped definitions
const (
	i         = 1e4
	MaxThread = 10
	prefix    = "George"
)

var (
	frenchHello string
	emptyString string = ""
)

func show_multiple_assignments() {
	fmt.Println("show_multiple_assignments()")
	var v1 int = 42

	var v2, v3 int = 2, 3

	fmt.Printf("vname1=%v, vname2=%v, vname3=%v\n", vname1, vname2, vname3)
	fmt.Printf("v1=%v, v2=%v, v3=%v\n", v1, v2, v3)
	fmt.Printf("b =", b)
}

func show_bool() {
	fmt.Println("show_bool()")
	var available bool
	valid := false
	available = true

	fmt.Printf("valid = %v, !valid = %v\n", valid, !valid)
	fmt.Printf("available=%v\n", available)
}

func show_different_types() {
	fmt.Println("Show_different_types()")
	var (
		unicodeChar rune
		a           int8
		b           int16
		c           int32
		d           int64
		e           byte
		f           uint8
	)
	var cmplx complex64 = 5 + 5i
	fmt.Println("Default values for int types")
	fmt.Println(unicodeChar, a, b, c, d, e, f)
	fmt.Printf("Value is: %v\n", cmplx)
}

func show_strings() {
	fmt.Println("show_strings()")
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Ohaiyou"
	frenchHello = "Bonjour"

	fmt.Println("Random String")
	fmt.Println(frenchHello, japaneseHello, no, yes, maybe)

	fmt.Println(`This is on multiple lines`)
}

func show_string_manipulation() {
	fmt.Println("Show_string_manipulation()")
	var s string = "hello"

	s = "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)

	m := " world"
	a := s + m

	d := "c" + s[1:]
	fmt.Printf("%s\n", d)
	fmt.Printf("s = %s, c=%v\n", s, c)
	fmt.Printf("s2=%s\n", s2)
	fmt.Printf("Combined strings\na = %s, d=%s\n", a, d)
}

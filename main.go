package main

import (
	"fmt"
	"reflect"
)

func main() {
	p1()
}

func p1() {
	f := 1.11
	numf := int(f)
	fmt.Println(numf)
	fmt.Println("numfの型は...", reflect.TypeOf(numf))

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}

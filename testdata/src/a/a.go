package a

import "fmt"

func f1(s []int) {
	s = append(s, 1) // want "this assignment is not detected outside of the func"
}

func f2(s *[]int) {
	*s = append(*s, 1) // OK
}

func f3(s []int) {
	var t []int
	t = append(s, 1) // OK
	fmt.Println(t)
}

func f4(s []int) {
	t := make([]int, 0)
	s = append(t, 1) // OK
}

func f5(s []int) {
	s, _ = append(s, 1), "foo" // want "this assignment is not detected outside of the func"
}

func f6(s []int) {
	s, _ = make([]int, 0), append(s, 1) // OK
}

func f7(s []int) {
	s = append(s, 1) // want "this assignment is not detected outside of the func"
	s = append(s, 2) // want "this assignment is not detected outside of the func"
}

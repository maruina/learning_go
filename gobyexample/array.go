package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("empty: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	s := make([]string, 3)
	fmt.Println("empty:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s: ", s)
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)
	s = append(s, "d")
	fmt.Println("appended: ", s)

	l := s[2:4]
	fmt.Println("sliced: ", l)

}

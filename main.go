package main

import "fmt"

type A struct {
	a int
}

type B struct {
	A
	b int
}

func main() {

	//c := NewController()
	//c.Run()
	s := []int{10, 20, 30}
	fmt.Println(s)

	s = append(s, 40)
	fmt.Println(s)

	s = append(s[0:2], s[3:]...)
	fmt.Println(s)

}

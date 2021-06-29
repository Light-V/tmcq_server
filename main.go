package main

type A struct {
	a int
}

type B struct {
	A
	b int
}

func main() {

	c := NewController()
	c.Run()

}

package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Speak() string
}

type MeatEater interface {
	Eat() string
	Animal
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Eat() string {
	return "EatMeat!"
}

func main() {

	c := NewController()
	c.Run()
	animal := Cat{}
	t := reflect.TypeOf(animal)

	fmt.Println(t)
}

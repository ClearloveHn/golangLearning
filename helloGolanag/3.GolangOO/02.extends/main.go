package main

import "fmt"

/**
golang使用结构体嵌套实现继承的概念
*/

type Animal struct {
	name string
	age  int
}

func (animal Animal) eat() {
	fmt.Println(animal.name + ",eat")
}

func (animal Animal) sleep() {
	fmt.Println(animal.name + ",sleep")
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func main() {
	dog := Dog{
		Animal{
			name: "dog",
			age:  2,
		},
	}

	cat := Cat{}
	cat.Animal.name = "cat"
	cat.age = 3

	dog.eat()
	dog.sleep()

	cat.eat()
	cat.sleep()
}

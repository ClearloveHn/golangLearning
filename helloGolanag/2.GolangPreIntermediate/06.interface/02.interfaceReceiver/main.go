package main

import "fmt"

/**
值接收者是一个拷贝,是一个副本,而指针接收者,传递的是指针。
*/

func main() {
	dog := Dog{
		"爱宝",
	}
	fmt.Printf("dog %p\n", &dog)
	dog.eat()
	fmt.Printf("dog %v\n", dog)

	fmt.Println("-------------")

	dogs := &Dog{
		"乐宝",
	}
	fmt.Printf("dog,%p\n", dogs)
	dogs.eats()
	fmt.Printf("dog %v\n", dogs)
}

type Pet interface {
	eat()
}

type Dog struct {
	name string
}

// 值接受者
func (dog Dog) eat() {
	fmt.Printf("dog: %p\n", &dog)
	fmt.Println("dog eat")
	dog.name = "福宝"
}

// 值接受者
func (dog *Dog) eats() {
	fmt.Printf("dog: %p\n", &dog)
	fmt.Println("dog eating")
	dog.name = "福宝"
}

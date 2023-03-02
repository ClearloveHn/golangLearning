package main

import "fmt"

/**
Go中的方法,是一种特殊的函数,定义于struct之上(与struct关联、绑定),被称为struct的接受者(receiver)。
*/

func main() {
	var person Person
	person.name = "messi"
	person.eat()
	person.sleep()
}

type Person struct {
	name string
}

func (person Person) eat() {
	fmt.Println(person.name + " eating")
}

func (person Person) sleep() {
	fmt.Println(person.name + " eating")
}

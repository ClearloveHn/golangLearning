package main

import "fmt"

/**
go结构体可以像普通变量一样，作为函数的参数，传递给函数，这里分为两种情况：
1.直接传递结构体，这是是一个副本（拷贝），在函数内部不会改变外面结构体内容。
2.传递结构体指针，这时在函数内部，能够改变外部结构体内容。
*/

func main() {
	person := Person{1, "tom"}
	fmt.Printf("person: %v\n", person)
	fmt.Println("=======")
	showPerson(person)
	fmt.Println("=======")
	fmt.Println(person)
}

type Person struct {
	id   int
	name string
}

func showPerson(person Person) {
	person.id = 1
	person.name = "son"
	fmt.Printf("person: %v\n", person)
}

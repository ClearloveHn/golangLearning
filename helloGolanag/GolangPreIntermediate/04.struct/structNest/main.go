package main

import "fmt"

/**
使用嵌套来实现继承
*/

type Dog struct {
	name  string
	color string
	age   int
}
type Person struct {
	dog  Dog
	name string
	age  int
}

func main() {
	var tom Person
	tom.dog.name = "花花"
	tom.dog.color = "黑白色"
	tom.dog.age = 2

	tom.name = "tom"
	tom.age = 20

	fmt.Printf("tom %v\n", tom)

}

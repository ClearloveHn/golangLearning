package main

import "fmt"

func main() {
	type Person struct {
		id   int
		name string
	}

	var tom = Person{1, "tom"}

	//go结构体指针
	var pPerson *Person
	pPerson = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", pPerson)
	fmt.Printf("*p_person: %v\n", *pPerson)

	//new关键字创建结构体指针
	var pPerson2 = new(Person)
	fmt.Printf("%T\n", pPerson2)

	//访问结构体指针成员
	pPerson.id = 1
	pPerson.name = "tom"
	fmt.Printf("%v\n", *pPerson)

}

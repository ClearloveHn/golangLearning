package main

import "fmt"

/**
golang没有面向对象的概念,也没有封装的概念,但是可以通过结构体struct和函数绑定来实现OOP的属性和方法等特性。接收者receiver方法。
*/

type Person struct {
	name string
	age  int
}

func (per Person) eat() {
	fmt.Println("eat...")
}

func (per Person) sleep() {
	fmt.Println("sleep...")
}

func (per Person) work() {
	fmt.Println("work...")
}

func main() {
	per := Person{
		name: "tom",
		age:  20,
	}

	fmt.Printf("per: %v\n", per)

	per.eat()
	per.sleep()
	per.work()
}

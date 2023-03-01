package main

import "fmt"

/**
接口注意事项：
1.go语言的接口,是一种新的类型定义,它把所有的具有共性的方法定义在一起,任何其他类型只要实现了这些方法就是实现了这个接口。
2.实现接口必须实现接口中所有的方法。
*/

func main() {
	customer := Customer{}
	customer.read()
	customer.write()

	mobile := Mobile{}
	mobile.read()
	mobile.write()
}

type USB interface {
	read()
	write()
}

type Customer struct {
}

type Mobile struct {
}

func (customer Customer) read() {
	fmt.Println("customer read")
}

func (customer Customer) write() {
	fmt.Println("customer write")
}

func (mobile Mobile) read() {
	fmt.Println("mobile read")
}

func (mobile Mobile) write() {
	fmt.Println("mobile write")
}

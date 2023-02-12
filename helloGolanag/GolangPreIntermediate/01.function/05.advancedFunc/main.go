package main

import "fmt"

/**
1.函数可以作为参数和返回值去使用
*/

func main() {
	test("lhn", sayHello)

	add := cal("+")
	r := add(1, 2)
	fmt.Println(r)

	sub := cal("-")
	s := sub(100, 50)
	fmt.Println(s)
}

func sayHello(name string) {
	fmt.Printf("Hello,%v\n", name)
}

// 函数作为参数
func test(name string, f func(string)) {
	f(name)
}

// 函数作为返回值
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

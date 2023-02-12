package main

import "fmt"

/**
1.闭包=函数+引用环境
2.闭包可以理解为定义在函数内部的函数,比如匿名函数,闭包是将函数内部和函数外部连接起来的桥梁。
*/

func main() {
	var f = add()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println("=======")
	f1 := add()
	fmt.Println(f1(40))
	fmt.Println(f1(50))
}

// 1.闭包实例1
func add() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

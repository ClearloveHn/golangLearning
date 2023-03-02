package main

import "fmt"

func main() {
	//调用函数-接收返回值
	sum := add(2, 3)
	fmt.Printf("%v", sum)
}

// 定义一个求和函数
func add(a int, b int) int {
	sum := a + b
	return sum
}

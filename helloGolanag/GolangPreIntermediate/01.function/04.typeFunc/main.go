package main

import "fmt"

func main() {
	var f fun
	f = sum
	s := f(1, 2)
	fmt.Printf("%v\n", s)

	f = max
	m := f(6, 7)
	fmt.Printf("%v", m)

}

// 使用type定义一个函数类型
type fun func(int, int) int

// 定义一个求和函数
func sum(a int, b int) int {
	return a + b
}

// 比较大小函数
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

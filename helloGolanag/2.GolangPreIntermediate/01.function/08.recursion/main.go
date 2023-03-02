package main

import "fmt"

/**
1.递归就是自己调用自己。
2.必须先定义函数的退出条件,没有退出条件,递归会成为死循环。
3.go语言递归函数很可能会产生一大堆的goroutine，也很可能会出现栈空间内存溢出问题。
*/

func main() {
	n := 5
	r := a(n)
	fmt.Println(r)
	fmt.Println("======")

	result := f(5)
	fmt.Println(result)
}

// 阶乘
func a(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * a(n-1)
	}
}

// 斐波那契数列
func f(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return f(n-1) + f(n-2)
}

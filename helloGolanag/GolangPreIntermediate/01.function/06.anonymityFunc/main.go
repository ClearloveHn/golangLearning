package main

import "fmt"

/**
1.go中函数不能嵌套,但是可以在函数内部定义匿名函数,实现一下简单功能调用。
2.匿名函数都是闭包
*/

func main() {
	func(a int, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Println(max)
	}(1, 2)

}

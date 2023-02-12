package main

import "fmt"

const max = 3 //定义一个常量

func main() {
	a := []int{1, 2, 3}
	var i int
	var ptr [max]*int
	fmt.Println(ptr)
	for i = 0; i < max; i++ {
		ptr[i] = &a[i] //整数地址赋值给指针数组
		fmt.Println(i, *ptr[i])
	}
}

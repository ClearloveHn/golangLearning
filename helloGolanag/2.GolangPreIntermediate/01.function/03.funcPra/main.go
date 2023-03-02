package main

import "fmt"

/**
1.声明函数时的参数列表叫做形参,调用时传递的参数叫做实参。
2.go语言是通过传值的方式传参的,意味着传递给函数的是拷贝后的副本,所以函数内部访问,修改的也是这个副本。
3.map、slice、interface、channel这些数据类型本身就是指针类型的,所以就算是拷贝传值也是拷贝的指针,
拷贝后的参数仍然指向底层数据结构,所以修改它们可能会影响外部数据结构的值。
*/

func main() {
	a := 100
	f1(a)
	fmt.Println(a)

	b := []int{1, 2}
	f2(b)
	fmt.Println(b)
}

// 参数传递(值传递)
func f1(a int) {
	a = 200
	fmt.Println(a)
}

// 注释第三条
func f2(b []int) {
	b[0] = 100
}

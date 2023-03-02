package main

import "fmt"

func main() {
	//无返回值的函数
	f1()

	//一个返回值的函数
	add(3, 3)

	//多个返回值,且在return中指定返回的内容
	f2()

	//多个返回值,返回值名称没有被使用
	f3()

	//return覆盖命名返回值,返回值名称没有被使用
	f4()
}

func f1() {
	fmt.Println("我没有返回值,只是一些计算")
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func f2() (name string, age int) {
	name = "昊楠"
	age = 26
	return name, age
}

func f3() (name string, age int) {
	name = "昊楠"
	age = 26
	return // 等价于return name, age
}

func f4() (name string, age int) {
	n := "昊楠"
	a := 26
	return n, a
}

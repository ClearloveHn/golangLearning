package main

import "fmt"

/**
变量注意事项:
1.变量一旦声明必须使用。
2。变量如果没有初始化,系统自动给默认值。
3。短声明只适合在函数内部使用。
*/

func main() {
	//声明一个变量
	var name string

	//声明变量并初始化
	var names = "昊楠"

	//短声明变量
	goal := 10

	//输出变量
	fmt.Println(name, names, goal)

}

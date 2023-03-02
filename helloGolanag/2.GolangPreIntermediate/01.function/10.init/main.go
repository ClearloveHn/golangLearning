package main

import "fmt"

/*
1.init函数先于main函数执行,不能被其他函数调用。
2.init函数没有参数和返回值。
3.每个包可以有多个init函数。
4.包的每个源文件也可以有多个init函数。
5.不同包的init函数按包导入的依赖关系决定执行顺序。
6.golang初始化顺序:变量初始化->init()->main()
*/

var a int = initVar()

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init")
}

func initVar() int {
	fmt.Println("init var...")
	return 100
}

func main() {
	fmt.Println("main...")
}

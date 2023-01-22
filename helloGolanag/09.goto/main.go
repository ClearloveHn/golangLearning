package main

import "fmt"

/**
goto用法
1.goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
2.Go语言中使用goto语句能简化一些代码的实现过程。例如双层嵌套的for循环要退出时。
*/

func main() {
	//test()
	test2()
}

// 跳转到指定标签
func test() {
	a := 0
	if a == 1 {
		goto LABEL1
	} else {
		fmt.Println("other")
	}

LABEL1:
	fmt.Printf("next...")
}

// 跳出双重for循环
func test2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto LABEL1
			}
		}
	}
LABEL1:
	fmt.Println("label1")
}

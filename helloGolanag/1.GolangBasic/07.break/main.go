package main

import "fmt"

/**
break注意事项
1.单独在select中使用break和不使用break没有啥区别。
2.单独在表达式switch语句，并且没有fallThough，使用break和不使用break没有啥区别。
3.单独在表达式switch语句，并且有fallThough，使用break能够终止fallThough后面的case语句的执行。
4.带标签的break,可以跳出多层select/switch作用域。让break更加灵活，写法更加简单灵活，不需要使用控制变量一层一层跳出循环，
没有带标签的break的只能跳出当前语句块。
*/

func main() {
	//test()
	test2()
}

// break和fallthrough一起使用
func test() {
	i := 2
	switch i {
	case 1:
		fmt.Println("等于1")
		break
	case 2:
		fmt.Println("等于2")
		fallthrough
	case 3:
		fmt.Println("等于3")
		break
	default:
		fmt.Println("不关心")
		break
	}
}

// break带标签
func test2() {
outLable:
	for i := 0; i < 5; i++ {
		fmt.Printf("外层：第%d次外层循环\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("内层：第%d次内层循环\n", j)
			//break
			break outLable
		}
		fmt.Printf("外层：没有跳过第%d次循环\n", i)
	}
}

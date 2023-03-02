package main

import "fmt"

func main() {
	test()
	test1()
	test2()
	test3()
}

func test() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}
}

// switch多条件匹配
func test1() {
	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	}
}

// case为条件表达式
func test2() {
	score := 90
	switch {
	case score >= 90:
		fmt.Println("享受假期")
	case score < 90 && score >= 80:
		fmt.Println("好好学习吧！")
	default:
		fmt.Println("玩命学习！")
	}
}

// fallthrough //执行多个case(满足当前条件的下一个case)
func test3() {
	a := 100
	switch a {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("other")
	}
}

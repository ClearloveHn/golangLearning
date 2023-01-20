package main

import "fmt"

func main() {
	test()
	test1()
	test2()
	test3()
}

// if语句
func test() {
	if age := 20; age > 18 {
		fmt.Println("你是成年人")
	}
}

// if-else语句
func test1() {
	if age := 17; age > 18 {
		fmt.Println("你是成年人")
	} else {
		fmt.Println("你是未成年人")
	}
}

// if-else-if
func test2() {
	if score := 80; score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score > 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

// if嵌套
func test3() {
	a := 100
	b := 200
	c := 3
	if a > b {
		if a > c {
			fmt.Println("a最大")
		}
	} else {
		if b > c {
			fmt.Println("b最大")
		} else {
			fmt.Println("c最大")
		}
	}
}

package main

import "fmt"

func main() {
	//test()
	//test1()
	test2()
}

func test() {
	//定义一个数组
	var a [3]int
	//定义一个数组并初始化
	var b = [4]int{1, 2, 3, 4}
	//短声明数组
	c := [1]string{"我爱你"}
	fmt.Printf("%v\n%v\n%s", a, b, c)
}

// 访问数组
func test1() {
	var a [2]int
	a[0] = 100
	a[1] = 200

	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Printf("a[0]: %v\n", a[1])

	// 修改 a[0]  a[1]
	a[0] = 1
	a[1] = 2

	fmt.Println("-----------")

	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Printf("a[0]: %v\n", a[1])

}

// for 循环遍历数组
func test2() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]: %v\n", i, a[i])
	}

}

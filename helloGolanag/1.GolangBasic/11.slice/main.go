package main

import "fmt"

func main() {
	//test()
	//test1()
	//test2()
	test3()
}

// 声明切片->初始化
func test() {
	var a []int
	b := make([]int, 10)
	fmt.Printf("%v\n%v", a, b)
}

// 切片添加元素
func test1() {
	var s1 []int
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3, 4, 5) // 添加多个元素
	fmt.Printf("s1: %v\n", s1)
}

// 切片删除元素
func test2() {
	s1 := []int{1, 2, 3, 4, 5}
	// 删除索引为2的元素
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

// copy
func test3() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s1[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("----------")

	s3 := make([]int, 3)
	copy(s3, s1)
	s1[0] = 1

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s3: %v\n", s3)

}

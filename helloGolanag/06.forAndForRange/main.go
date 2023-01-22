package main

import "fmt"

/**
for range注意事项
1.Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel)。
2.数组、切片、字符串返回索引和值。
3.map返回键和值。
4.通道（channel）只返回通道内的值。
*/

func main() {
	//test()
	//test2()
	//test3()
	test4()
}

// 普通for循环
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// for range -数组
func test2() {
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i: %d, v: %v\n", i, v)
	}
}

// for range -切片
func test3() {
	var a = []int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i: %d, v:%d\n", i, v)
	}
}

// for range map
func test4() {
	m := make(map[string]string)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "tom@gmail.com"
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
}

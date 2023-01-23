package main

import "fmt"

func main() {
	//test()
	//test2()
	test3()
}

// 声明map并且初始化
func test() {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"
	fmt.Printf("m1: %v\n", m1)

	m2 := map[string]string{
		"name":  "kite",
		"age":   "20",
		"email": "kite@gmail.com",
	}
	fmt.Printf("m2: %v\n", m2)
}

// 根据key访问map
func test2() {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	name := m1["name"]
	age := m1["age"]
	email := m1["email"]
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
}

// 判断某个key是否存在
func test3() {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	v, ok := m1["address"]
	if ok {
		fmt.Println("键存在")
		fmt.Println(v)
	} else {
		fmt.Println("键不存在")
	}
}

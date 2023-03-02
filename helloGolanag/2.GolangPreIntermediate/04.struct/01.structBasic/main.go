package main

import "fmt"

func main() {

	//定义一个person结构体
	type Person struct {
		id, age     int //成员变量
		name, email string
	}

	//声明一个结构体变量
	var tom Person
	fmt.Println(tom)

	//使用.访问结构体变量
	var son Person
	son.id = 0
	son.name = "son"
	son.age = 25
	son.email = "xxx"
	fmt.Printf("%v\n", son)

	//匿名结构体

}

package main

import "fmt"

func main() {

	//使用键值对,对结构体进行初始化。
	type Person struct {
		id, age   int
		name, sex string
	}
	kite := Person{
		id:   1,
		name: "kite",
		age:  20,
		sex:  "男",
	}
	fmt.Printf("%v", kite)

	//使用值的列表初始化
	/**注意事项
	1.必须初始化结构体的所有字段。
	2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	3.该方式不能和键值初始化方式混用。
	*/
	hat := Person{
		1,
		20,
		"hat",
		"男",
	}
	fmt.Printf("kite: %v\n", hat)

}

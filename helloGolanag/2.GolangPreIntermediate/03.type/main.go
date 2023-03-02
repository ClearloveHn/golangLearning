package main

import "fmt"

/**
类型定义和类型别名注意事项:
区别:
	1.类型定义相当于定义了一个全新的类型，与之前的类型不同；但是类型别名并没有定义一个新的类型，而是使用一个别名来替换之前的类型。
    2.类型别名只会在代码中存在，在编译完成之后并不会存在该别名。
    3.因为类型别名和原来的类型是一致的，所以原来类型所拥有的方法，类型别名中也可以调用，但是如果是重新定义的一个类型，那么不可以调用之前的任何方法。
*/

func main() {
	//类型定义
	type myInt int
	var i myInt // i 为myInt类型
	i = 100
	fmt.Printf("i: %v i: %T\n", i, i)

	//类型别名
	type myInt2 = int
	var j myInt2 //实际j还是myInt2类型
	fmt.Printf("j: %v j: %T\n", j, j)
}

package main

/*
*
常量注意事项:
1.常量可以是多种类型
2.iota是个特殊的常量(可以被编译器修改的常量)。
3.iota默认值是0,每调用一次便+1,可以插队,遇到const便重置为0。
*/

func main() {

	//声明一个常量
	const pi float64 = 3.14

	//iota(特殊常量,可以插队)
	const (
		a = iota
		b = iota
		c = iota
	)

	//iota插队
	const (
		a1 = iota
		b1 = 100
		c1 = iota
	)

}

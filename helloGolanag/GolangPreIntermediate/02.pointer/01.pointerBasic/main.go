package main

import "fmt"

/*
*
1.&取址,*取值
*/
func main() {
	var a int = 20 //声明实际变量
	var ip *int    // 声明指针变量
	ip = &a        //指针变量的存储地址
	fmt.Println("a变量的地址是:", &a)
	fmt.Println("ip变量储存的指针地址是", ip)
	fmt.Println("ip变量的值是", *ip)
}

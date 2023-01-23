package main

import "fmt"

func main() {

	str := "hello world"
	n := 3
	m := 5
	fmt.Println(str[n])   //获取字符串索引位置为n的原始字节
	fmt.Println(str[n:m]) //截取得字符串索引位置为 n 到 m-1 的字符串
	fmt.Println(str[n:])  //截取得字符串索引位置为 n 到 len(s)-1 的字符串
	fmt.Println(str[:m])  //截取得字符串索引位置为 0 到 m-1 的字符串

}

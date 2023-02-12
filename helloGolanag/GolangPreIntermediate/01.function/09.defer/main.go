package main

import "fmt"

/**
1。defer语句会将其后面跟随的语句进行延迟处理。
2。多个defer语句,按照先进后出的方式执行。
*/

func main() {
	fmt.Println("start")
	defer fmt.Printf("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")
}

package main

/*
Golang中的并发是函数相互独立运行的能力。Goroutines是并发运行的函数。Golang提供了Goroutines作为并发处理操作的一种方式。
*/

import (
	"fmt"
	"time"
)

// 实例1
func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg:%v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() { //main函数调用的是主协程
	go show("java") //启动了一个协程来执行
	show("golang")
	fmt.Println("end")
}

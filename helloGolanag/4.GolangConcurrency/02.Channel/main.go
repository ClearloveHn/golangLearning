package main

import "math/rand"

/**
channel注意事项：
1.Go 提供了一种称为channel的机制，用于在 goroutine 之间共享数据。
2.当使用goroutine执行并发活动时，需要在goroutine之间共享资源或数据，channel充当goroutine之间的管道，并提供一种机制来保证同步交换。
3.数据在channel上传递：在任何给定时间只有一个goroutine可以访问数据项，因此不会发生数据竞争。
4.有两种类型的channel：无缓冲channel和有缓冲channel。
5.无缓冲channel:
	1.用于同步。
	2.接收到任何值之前没有能力保存它。
	3.发送和接收goroutine在任何发送或接收操作完成之前的同一时刻都准备就绪。
    如果两个goroutine没有在同一时刻准备好,则channel会让执行其各自发送或接收操作的goroutine首先等待。
6.有缓冲channel:
	1.用于异步。
	2.有能力在接收到一个或多个值之前保存它们。
	3.只有当channel中没有要接收的值时,接收才会阻塞。仅当没有可用缓冲区来放置正在发送的值时,发送才会阻塞。
*/

import (
	"fmt"
	"time"
)

// 创建int类型channel
var valuesChannel = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	valuesChannel <- value
}

func main() {
	//从channel接收值
	defer close(valuesChannel)
	go send()
	fmt.Println("wait")
	value := <-valuesChannel
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")
}

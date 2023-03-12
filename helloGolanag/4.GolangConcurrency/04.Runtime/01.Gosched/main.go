package main

import (
	"fmt"
	"runtime"
)

/**
让出cpu时间片,重新等待安排任务。类似于waitGroup用法。
*/

func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}

}
func main() {
	go show("java")
	for i := 0; i < 2; i++ {
		runtime.Gosched() //cpu切一下,再次分配任务。
		fmt.Println("golang")
	}
}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func show() {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit() //ιεΊεη¨
		}
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	go show()
	time.Sleep(time.Second)

}

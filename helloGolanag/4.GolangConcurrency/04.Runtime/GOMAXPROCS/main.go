package main

import (
	"fmt"
	"runtime"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) //查看cpu核心数
	//runtime.GOMAXPROCS(1)                                  //cpu核心数设置为1
	go a()
	go b()
	runtime.Gosched()
}

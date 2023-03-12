package main

import (
	"fmt"
	"sync"
)

/**
除了使用channel实现同步之外,还可以使用Mutex互斥锁的方式实现同步。
*/

var m int = 100

var lock sync.Mutex

var wt sync.WaitGroup

func add() {
	defer wt.Done()
	lock.Lock()
	m += 1
	lock.Unlock()
}

func sub() {
	defer wt.Done()
	lock.Lock()
	m -= 1
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		wt.Add(1)
		go sub()
		wt.Add(1)
	}
	wt.Wait()
	fmt.Printf("m: %v\n", m)
}

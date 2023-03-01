package main

import "fmt"

/**
接口可以通过嵌套，创建新的接口。
*/

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}
type Fish struct {
}

func (fish Fish) fly() {
	fmt.Println("fly.......")
}

func (fish Fish) swim() {
	fmt.Println("swim.......")
}

func main() {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}

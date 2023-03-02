package main

import "fmt"

/**
1.一个类型可以实现多个接口。
2.多个类型可以实现同一个接口(多态)。
*/

func main() {
	//一个类型实现多个接口
	mobile := Mobile{}
	mobile.playMusic()
	mobile.playVideo()

	//多个类型实现同一个接口(多态)
	var p Pet
	p = Cat{}
	p.eat()
	p = Dog{}
	p.eat()
}

// 一个类型实现多个接口

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct {
}

func (mobile Mobile) playMusic() {
	fmt.Println("播放音乐")
}

func (mobile Mobile) playVideo() {
	fmt.Println("播放视频")
}

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat")
}

func (cat Cat) eat() {
	fmt.Println("cat eat")
}

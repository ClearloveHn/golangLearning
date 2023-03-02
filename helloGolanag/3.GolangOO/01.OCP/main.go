package main

import "fmt"

/**
1.使用接口的这种设计方法，可以很好的解耦合代码，实现软件设计的OCP原则（即开闭原则）。
2.这样设计，如果再添加一个宠物，例如：一个鸟Bird，原有的代码不用修改，直接添加就可以。
*/

// Pet 宠物接口pet
type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eat")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep")
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat")
}

func (cat Cat) sleep() {
	fmt.Println("cat sleep")
}

type Person struct {
	name string
}

// 为person添加一个养宠物的方法
func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{
		"福宝",
		4,
	}

	cat := Cat{
		"乐宝",
		8,
	}

	per := Person{
		"崔爷爷",
	}

	per.care(dog)
	per.care(cat)
}

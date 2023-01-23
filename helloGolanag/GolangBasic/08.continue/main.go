package main

import "fmt"

/**
continue的作用:
1.continue的作用是结束本次循环，即跳过循环体中下面尚未执行的语句，然后进行下一次是否执行循环的判定。
2.continue只能在for循环中使用
*/

func main() {
	//test1()
	test2()
}

// continue的使用
func test1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "能被2整除")
		} else {
			continue
		}
	}
}

func test2() {
	// MY_LABEL:
	for i := 0; i < 5; i++ {
	MyLabel:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue MyLabel
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}
}

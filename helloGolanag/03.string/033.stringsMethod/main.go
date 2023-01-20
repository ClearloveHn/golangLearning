package main

import (
	"fmt"
	"strings"
)

// strings常用方法:
func main() {
	s := "hello world！"

	//求长度
	fmt.Printf("len(s): %v\n", len(s))

	//分割
	fmt.Printf("strings.Split(s, \"\"): %v\n", strings.Split(s, " "))

	// 字符串是否包含
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))

	//前/后缀是否存在(判断)
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("strings.HasSuffix(s, \"world！\"): %v\n", strings.HasSuffix(s, "world！"))

	//字符出现的位置
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))
}

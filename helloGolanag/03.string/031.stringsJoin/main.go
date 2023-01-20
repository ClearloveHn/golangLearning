package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//字符串拼接(1)
	var buffer bytes.Buffer
	buffer.WriteString("lhn")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Println(buffer.String())

	//字符串拼接(2)
	s1 := "lhn"
	s2 := "21"
	msg := fmt.Sprintf("%s,%s", s1, s2)
	fmt.Println(msg)

	//字符串拼接(3)
	name := "lhn"
	age := "22"
	msg2 := strings.Join([]string{name, age}, ",")
	fmt.Println(msg2)

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	practices()
	printName()
}

func practices() {
	// 统计一下文章中单词 better 出现的次数，并进行输出。
	// Beautiful is better than ugly.
	// Explicit is better than implicit.
	// Simple is better than complex.
	p := `Beautiful is better than ugly.Explicit is better than implicit.Simple is better than complex.`
	fmt.Printf("count: %d\n", strings.Count(p, "better"))

}

func printName() {
	// 请编写程序，在程序运行后提示用户输入自己的名字，并将名字的首字母改为大写并输出 Hello 用户名
	fmt.Print("input your name: ")
	var name string
	fmt.Scanln(&name)                             // 获取用户输入的值
	join := strings.ToUpper(name[0:1]) + name[1:] // 首字母大写
	fmt.Printf("hello %s\n", join)
}

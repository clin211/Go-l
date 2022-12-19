package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	practices()
	printName()
	handleMatchString()
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

// 使用 Regex 验证字符串是否包含空格
func handleMatchString() {
	word := "go language"

	// regexp.MustCompile() 函数用于创建正则表达式，MatchString()函数返回一个布尔值，指示模式是否与字符串匹配
	whitespace := regexp.MustCompile(`\s`).MatchString(word)
	println("is space: ", whitespace) // true
}

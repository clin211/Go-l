package main

import (
	"fmt"
)

// 三总声明变量的方式
// 1、显式为变量赋值： var name string = "clin"
// 2、短变量的方式：name := "clin"
// 3、变量声明快： var (
//
//		name string
//		age uint
//		isWorking bool
//	)
func main() {
	var name string = "clin"
	fmt.Print(name)

	nickname := "forest"
	fmt.Print(nickname)

	var (
		age       uint
		isWorking bool
	)

	age = 23
	isWorking = true
	fmt.Print(age)
	fmt.Print(isWorking)
}

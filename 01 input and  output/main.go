package main

import "fmt"

func main() {
	//可以从控制台接收用户信息，【姓名，年龄，薪水, 是否通过考试 】。
	//先声明所需的变量
	var name string
	var age byte
	var sal float32
	fmt.Println("input your name: ") // 当程序执行到本行时，会阻塞到这里，等待用户输入
	fmt.Scanln(&name)

	fmt.Println("input your age: ")
	fmt.Scanln(&age)

	fmt.Println("input your sal: ")
	fmt.Scanln(&sal)

	fmt.Printf("name: %v \nage: %v \nsal: %v \n",
		name, age, sal)
}

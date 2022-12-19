package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	ConvertStringsToValuesUsingAtoiType()

	fmt.Println("-------------------------------")
	handleParseInt()

	fmt.Println("-------------------------------")
	useSscanConvertDataType()
}

// 使用 Atoi 将字符串转换为整数类型
func ConvertStringsToValuesUsingAtoiType() {
	str := "100"

	// Atoi()：Atoi代表ASCII转整数，Atoi()函数返回两个值：转换后的结果和错误（如果有）
	intVar, err := strconv.Atoi(str)
	if err != nil {
		println("convert error: ", err)
	}
	fmt.Println(intVar)                              // 100
	fmt.Println(intVar, err, reflect.TypeOf(intVar)) // 100 <nil> int
}

// 使用ParseInt()转换类型
func handleParseInt() {
	strVar := "100"

	// ParseInt() 以给定的基数（0、2到36）和位大小（0到64）解释字符串s并返回响应的值；该函数接受一个字符串参数，根据一个基本参数将其转换成对应的int类型；默认情况下是Int64值
	intVar, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar)) // 100 <nil> int64

	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar)) // 100 <nil> int64

	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar)) // 100 <nil> int64

	intVar, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar)) // 100 <nil> int64
}

// 使用Sscan方法
func useSscanConvertDataType() {
	str := "100"
	intValue := 0
	_, err := fmt.Sscan(str, &intValue)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(intValue, err, reflect.TypeOf(intValue)) // 100 <nil> int
}

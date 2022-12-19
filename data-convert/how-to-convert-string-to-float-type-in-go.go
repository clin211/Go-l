package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	useParseFloatConvertDataType()
}

func useParseFloatConvertDataType() {
	// ParseFloat() 将字符串s转换为浮点数，精度有bit size指定：32表示float32,64表示float64，当bit size为32时，结果仍然是float64类型，但他可以在不改变其值的情况下转为float32

	s := "3.1415926"
	f, err := strconv.ParseFloat(s, 8)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(f, err, reflect.TypeOf(f)) // 3.1415926 <nil> float64

	s1 := "-3.141"
	f1, err := strconv.ParseFloat(s1, 8)
	fmt.Println(f1, err, reflect.TypeOf(f1)) // -3.141 <nil> float64

	s2 := "A-3141X"
	f2, err := strconv.ParseFloat(s2, 32)
	if err != nil {
		fmt.Println(err) // strconv.ParseFloat: parsing "A-3141X": invalid syntax
	} else {
		fmt.Println(f2, err, reflect.TypeOf(f2))
	}
}

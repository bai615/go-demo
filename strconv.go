package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// 基本数值转换
	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.Atoi("711"))

	// 解析
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14",32))

	// 格式化
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 8))
}
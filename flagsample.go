package main

import (
	"fmt"
	//"os"
	"flag"
)

func style() {
	// 格式化定义
	methodPtr := flag.String("method", "default", "method of sample")
	valuePtr := flag.Int("value", -1, "value of sample")

	// 解析
	flag.Parse()

	fmt.Println(*methodPtr, *valuePtr)

	/**

	D:\codes\go>flagsample.exe
	default -1

	D:\codes\go>flagsample.exe -method submit -value 123
	submit 123
	 */
}

func style2() {
	var method string
	var value int

	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")

	flag.Parse()

	fmt.Println(method, value)
}
func main() {

	/**
	fmt.Println(os.Args)
	D:\codes\go>flagsample.exe a b c
    [flagsample.exe a b c]
	 */

	style2()
}

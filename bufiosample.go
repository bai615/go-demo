package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func main() {
	strReader := strings.NewReader("Hello World")

	bufReader := bufio.NewReader(strReader)

	// Peek : 返回缓存的一个切片，该切片引用缓存中前n字节数据
	data, _ := bufReader.Peek(5)

	fmt.Println(data)
	fmt.Println(string(data))
	fmt.Println(bufReader.Buffered()) // 查看缓存多少个字符

	str,_ := bufReader.ReadString(' ') // 从bufReader中读取，读到空格时为止

	fmt.Println(str, bufReader.Buffered())

	///////////////////////////////////////////////////////////////////////////////////

	w := bufio.NewWriter(os.Stdout)

	fmt.Fprint(w,"Hello ")
	fmt.Fprint(w,"World")
	w.Flush() // 提交到输出设备
}

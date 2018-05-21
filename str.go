package main

import ( 
    "fmt"
    "strings"
)

func main() {

    /*
    // 声明一个string类型变量并赋值
    var str1 string = "\\\"" 
    
    // 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
    fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 。\n", str1)
    */

    s := "Hello World"

    // 是否包含
    fmt.Println(strings.Contains(s,"Hello"),strings.Contains(s,"?"))
    // 索引
    fmt.Println(strings.Index(s, "o"))

    ss := "1#2#345"
    // 分割字符串
    splitedStr := strings.Split(ss, "#")
    fmt.Println(splitedStr)

    // 合并字符串
    fmt.Println(strings.Join(splitedStr, "#"))

    // 前缀、后缀
    fmt.Println(strings.HasPrefix(s, "He"), strings.HasSuffix(s, "ld"))
}
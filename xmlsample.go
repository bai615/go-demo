package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type person1 struct {
	Name string `xml:",attr"` // 属性
	Age  int
}

func main() {
	p := person{Name: "davy", Age: 18}
	p1 := person1{Name: "tom", Age: 20}

	// 序列化
	// xml.Marshal(p) 输出为一个字符串
	// xml.MarshalIndent(p, "", " ") 输出为一个格式化的字符串
	if data, err := xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(data))
	}

	///////////////////////////////////////////////////////////////////////////////////////////

	var data []byte
	var err error

	// 序列化
	if data, err = xml.MarshalIndent(p1, "", " "); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	/////////////////////////////////////////////////////////////////////////////////////////////

	p2 := new(person)

	// 反序列化
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p2)

}

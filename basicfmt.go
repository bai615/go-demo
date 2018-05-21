package main

import (
	"fmt"
	"os"
)

type Data struct {
}

func (self Data) String() string {
	return "data"
}

func main() {

	fmt.Print("Hello World\n")
	fmt.Println("Hello World")


	fmt.Printf("num %v\n", "h")

	str := fmt.Sprintf("float %f", 3.14159)
	fmt.Print(str)

	fmt.Fprintln(os.Stdout, "A\n")

	fmt.Printf("%v\n",Data{})
}

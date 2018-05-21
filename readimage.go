package main

import (
	"os"
	"encoding/binary"
	"fmt"
)

func main() {
	file, _ := os.Open("image.bmp")
	defer file.Close()

	var headA, headB byte
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	var reservedA,reservedB uint16
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	fmt.Printf("%c %c", headA, headB)

	fmt.Println(headA,headB,size,reservedA,reservedB,offbits)
}

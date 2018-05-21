package main

import (
	"os"
	"encoding/binary"
	"fmt"
)

type BitmapInfoHeader struct{
	Size uint32
	Width int32
	Height int32
	Places uint16
	BitCount uint16
	Compression uint32
	SizeImage uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed uint32
	ClrImportant uint32
}

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


	//实例化
	infoHeader := new(BitmapInfoHeader)
	binary.Read(file,binary.LittleEndian,infoHeader)

	fmt.Println(infoHeader)
}

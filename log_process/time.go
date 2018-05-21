package main

import (
	"time"
	"log"
	"fmt"
)

func main(){
	//nginx time format
	nginx_time :="03/Apr/2017:08:29:05 +0800"
	t,err:=time.Parse("02/Jan/2006:15:04:05 -0700",nginx_time)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
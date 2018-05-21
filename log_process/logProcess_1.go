package main

/**
import (
	"strings"
	"fmt"
	"time"
)

type LogProcess struct {
	rc          chan string
	wc          chan string
	path        string // 读取文件的路径
	influxDBDsn string // influx data source
}

// 读取模块
func (l *LogProcess) ReadFromFile() {
	line := "message"
	l.rc <- line
}

// 解析模块
func (l *LogProcess) Process() {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

// 写入模块
func (l *LogProcess) WriteToInfluxDB() {
	fmt.Println(<-l.wc)
}

func main() {
	lp := &LogProcess{
		rc:          make(chan string),
		wc:          make(chan string),
		path:        "/tmp/access.log",
		influxDBDsn: "username&password..",
	}

	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToInfluxDB()

	time.Sleep(1 * time.Second)
}
**/
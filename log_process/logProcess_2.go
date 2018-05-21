package main
/*
import (
	"strings"
	"fmt"
	"time"
	"os"
	"bufio"
	"io"
)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  Reader // 读取器
	write Writer // 写入器
}

type ReadFromFile struct {
	path string // 读取文件的路径
}

type WriteToInfluxDB struct {
	influxDBDsn string // influx data source
}

// 读取模块
func (r *ReadFromFile) Read(rc chan []byte) {
	//line := "message"
	//rc <- line

	// 打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}

	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2) // 文件指针移动到最后
	rd := bufio.NewReader(f)

	for {
		// 逐行读取
		line, err := rd.ReadBytes('\n')

		if err == io.EOF {
			// 读取到文件末尾，等待文件生成
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		//rc <- line
		rc <- line[:len(line)-1]//去掉换行符
	}

}

// 写入模块
func (w WriteToInfluxDB) Write(wc chan string) {
	//fmt.Println(<-wc)
	for v := range wc {
		fmt.Println(v)
	}
}

// 解析模块
func (l *LogProcess) Process() {
	//data := <-l.rc
	//l.wc <- strings.ToUpper(string(data))
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

func main() {
	r := &ReadFromFile{
		path: "./access.log",
	}

	w := &WriteToInfluxDB{
		influxDBDsn: "username&password..",
	}

	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(30 * time.Second)
}
*/
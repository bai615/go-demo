package main

import (
	"time"
	"fmt"
	"regexp"
)

func main()  {
	loc, _ := time.LoadLocation("Europe/Berlin")
	fmt.Println(loc)
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)
	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

	loc, _ = time.LoadLocation("Asia/Shanghai")
	//loc, _ = time.LoadLocation("Local")
	fmt.Println("----------------------------------------------------")
	fmt.Println(loc)
	t,_ = time.ParseInLocation("2006-Jan-02", "08/Jan/2016:10:40:18 +0800",loc)
	fmt.Println(t)

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//获取本地location
	toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ = time.LoadLocation("Local")                            //重要：获取时区
	fmt.Println(loc)
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)

	fmt.Println("=========================================")
	layout := "01__02-2006 3.04.05 PM"
	layout = "02 Jan 06 15:04 -0700"

	fmt.Println(time.Now().Format(layout))

	loc, _ = time.LoadLocation("Asia/Shanghai")
	fmt.Println(loc)
	t,_ = time.ParseInLocation("02/Jan/2006:15:04:05 +0000", "02/Jan/2016:10:40:18 +0800",loc)
	fmt.Println(t)

	fmt.Println(time.Now())

	fmt.Println("====================================================")
	log := "127.0.0.1 - - [23/Mar/2018:13:06:10 +0800] \"GET /debug/default/toolbar?tag=5ab48b4229538 HTTP/1.1\" 404 1214";
	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)
	ret := r.FindStringSubmatch(string(log)) //匹配数据内容
	fmt.Println(ret)
}

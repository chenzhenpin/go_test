package main

import (
	"fmt"
	"time"
)

func main()  {
	//sec, nsec := now()
	//第一个参数字符串2006-01-02 15:04:05代表格式,相当于java的yyyy-MM-dd HH:mm:ss字符串。
	//t, _ := time.Parse("2006-01-02 15:04:05", "2016-06-13 09:14:00")
	//ti,_ :=time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 09:14:00",time.Local)
	//fmt.Println(time.Now().Sub(t).Hours())
	//fmt.Println(time.Now().Sub(ti).Hours())
	//tt:=time.Now()
	//fmt.Println(tt.Format("2006-01-02 15:04:05"))
	//fmt.Println(tt.Format("2006年01月02日 15时04分05秒"))//必须使用2006-01-02 15:04:05这几数字进行格式化。
	ts,_ :=time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 09:39:50",time.Local)

	tp,_ :=time.ParseInLocation("2006-01-02 15:04:05", "2019-01-02 11:49:59",time.Local)
	fmt.Println(ts.Format("2006-01-02 15:04:05"))
	// 整点（向下取整）
	fmt.Println(ts.Truncate(1 * time.Hour))
	// 整点（最接近）
	fmt.Println(ts.Round(1 * time.Hour))

	// 整分（向下取整）
	fmt.Println(ts.Truncate(1 * time.Minute))
	// 整分（最接近）
	fmt.Println(ts.Round(1 * time.Minute))

	fmt.Println(tp.Sub(ts).Seconds())

}

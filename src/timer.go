package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})
	time.AfterFunc(1500*time.Millisecond, func() {
		fmt.Println("after func 1Second callback, elaspe:", time.Now().Sub(start))
	})

	time.Sleep(3 * time.Second)
	fmt.Println("sleep 1 Second")
	// time.Sleep(3*time.Second)
	// Reset 在 Timer 还未触发时返回 true；触发了或Stop了，返回false
	if timer.Reset(3 * time.Second) {//Reset 会先调用 stopTimer 再调用 startTimer，类似于废弃之前的定时器，重新启动一个定时器。返回值和 Stop 一样。
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}


	time.Sleep(10 * time.Second)
}

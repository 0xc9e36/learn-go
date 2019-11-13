/**
 * @Author: wei.tan
 * @Description:
 * @File:  timer
 * @Version: 1.0.0
 * @Date: 2019-10-28 21:51
 */

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//delayTime()
	AfterFuncDemo()
}

func AfterFuncDemo() {
	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(1 * time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})

	time.Sleep(2 * time.Second) // 等待协程退出
}

func delayTime()  {
	timer := time.NewTimer(6 * time.Second)

	timer.Stop()

	select {
	case <- timer.C:
		fmt.Println("延迟执行")
	}
}
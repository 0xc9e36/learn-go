/**
 * @Author: wei.tan
 * @Description:
 * @File:  ticker
 * @Version: 1.0.0
 * @Date: 2019-10-28 22:15
 */

package main

import (
	"log"
	"time"
)

// TickerDemo 用于演示ticker基础用法
func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

func main() {
	//TickerDemo()
	//fmt.Println(time.Now())
	time.Now().Round(0)
}
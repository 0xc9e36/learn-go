/**
 * @Author: wei.tan
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2019-11-13 21:33
 */

package main

import (
	"fmt"
	"reflect"
)

type Server struct {
	ServerName string `key1:"value1" key11:"value11"`
	ServerIP   string `key2:"value2"`
}


//反射获取 tag
func main() {
	s := Server{}
	st := reflect.TypeOf(s)

	field1 := st.Field(0)
	fmt.Printf("key1:%v\n", field1.Tag.Get("key1"))
	fmt.Printf("key11:%v\n", field1.Tag.Get("key11"))

	filed2 := st.Field(1)
	fmt.Printf("key2:%v\n", filed2.Tag.Get("key2"))
}
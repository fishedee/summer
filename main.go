package main

import (
	"fmt"
	"reflect"
)

type MM2 interface {
	Run()
}
type MM struct {
}

func (this *MM) Run() {

}

type UserAo interface {
	Get(id int) int
}

func do(a int) int {
	return a + 1
}

func mm(a ...interface{}) int {
	return a[0].(int) + 1
}

func main() {
	var k interface{}
	k = mm
	fmt.Println(reflect.ValueOf(k).Type().NumIn())
	fmt.Println(reflect.ValueOf((*MM2)(nil)).Type().Elem().Name())
	fmt.Println(reflect.ValueOf((*MM2)(nil)).Type().Elem().Method(0).Name)
	fmt.Println(reflect.ValueOf(MM2.Run).Type().In(0))
	fmt.Println(reflect.ValueOf(MM2.Run).Type().In(0))
}

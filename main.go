package main

import (
	"fmt"
	"reflect"
)

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
}

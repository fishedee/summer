package main

import (
	"fmt"
	"reflect"
)

type UserAo interface {
	Get(id int) int
}

func do(v interface{}) {
	a := reflect.ValueOf(v)
	fmt.Println(a)
	fmt.Println(a.Type())
}

func main() {
	do((UserAo).Get)
}

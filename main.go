package main

import (
	"fmt"
	"reflect"
	"runtime"
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

func main() {
	targetType := reflect.TypeOf(do)
	mm := reflect.MakeFunc(targetType, func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{reflect.ValueOf(3)}
	})
	dd := (mm.Interface()).((func(int) int))
	fmt.Println(dd(34))
	fmt.Println(targetType.Name())
	//mc := &MM{}
	fmt.Println(runtime.FuncForPC(mm.Pointer()).Name())
}

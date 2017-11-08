package main

import (
	"fmt"
	"github.com/fishedee/summer/core"
	"github.com/fishedee/summer/core_test/service"
)

func main() {
	context := core.NewContext()
	fmt.Println("------1------")
	result := service.DoService(context, 1, 2)
	fmt.Println(result)

	fmt.Println("------2------")
	context.PreHook(func() {
		fmt.Println("preHook")
	})
	context.PostHook(func() {
		fmt.Println("postHook")
	})
	result2 := service.DoService(context, 1, 2)
	fmt.Println(result2)

	fmt.Println("------3------")
	context.Mock(&service.DoService, func(context core.Context, a int, b int) int {
		fmt.Println(a, "*", b)
		return a * b
	})
	result3 := service.DoService(context, 1, 2)
	fmt.Println(result3)
}

package service

import (
	"fmt"
	"github.com/fishedee/summer/core"
)

func doServiceImplement(context core.Context, a int, b int) int {
	fmt.Println(a, "+", b)
	return a + b
}

func init() {
	core.RegisterBean(&DoService, doServiceImplement)
}

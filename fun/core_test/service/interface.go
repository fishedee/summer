package service

import (
	"github.com/fishedee/summer/core"
)

var DoService func(context core.Context, a int, b int) int

func init() {
	core.RegisterInterface(&DoService)
}

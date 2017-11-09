package main

import (
	"fmt"
	"github.com/fishedee/summer/ioc"
	"github.com/fishedee/summer/sample/api"
	_ "github.com/fishedee/summer/sample/service"
)

type MM struct {
	userAo api.UserAo
}

func (this *MM) run() {
	this.userAo.Add(api.User{
		Name: "123",
	})
	this.userAo.Add(api.User{
		Name: "456",
	})
	fmt.Println(this.userAo.Get(10001))
	fmt.Println(this.userAo.Get(10002))
}

func NewMM(userAo api.UserAo) *MM {
	this := &MM{}
	this.userAo = userAo
	return this
}

func main() {
	mm := ioc.New((*MM)(nil)).(*MM)
	mm.run()
}

func init() {
	ioc.Register(NewMM)
}

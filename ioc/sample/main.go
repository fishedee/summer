package main

type MM struct {
	userAo UserAo
}

func (this *MM) run() {
	this.userAo.Add(User{
		Name: "123",
	})
	this.userAo.Add(User{
		Name: "456",
	})
	fmt.Println(this.userAo.Get(10001))
	fmt.Println(this.userAo.Get(10002))
}

func main() {
	mm := &MM{}
	ctx := ioc.Inject(mm)
	mm.Run()

	ctx.Mock(UserAo(nil), func(name string, args []reflect.Value) reflect.Value {

	})
	mm.Run()

	ctx.Stub(UserAo(nil), func(name string, args []reflect.Value) reflect.Value {

	})
	mm.Run()
}

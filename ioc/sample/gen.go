
type UserAoMock interface {
	Get(id int) User
	Add(data User) int
}

func init(){
	ioc.RegisterMock(UserAo(nil),&UserAoMock{})
}
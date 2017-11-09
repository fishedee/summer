package service

import (
	"github.com/fishedee/summer/ioc"
	"github.com/fishedee/summer/sample/api"
	"testing"
)

type UserDbStub struct {
}

func (this *UserDbStub) Get(id int) api.User {
	return api.User{Id: 10001, Name: "Fish"}
}

func (this *UserDbStub) Add(data api.User) int {
	return 10002
}

func TestUserAoGet(t *testing.T) {
	userAo := ioc.New((*api.UserAo)(nil), func() api.UserDb {
		return &UserDbStub{}
	}).(api.UserAo)
	left := userAo.Get(0)
	right := api.User{Id: 10001, Name: "Fish"}
	if left.Id != right.Id ||
		left.Name != right.Name {
		t.Errorf("Error!")
	}
}

func TestUserAoAdd(t *testing.T) {
	userAo := ioc.New((*api.UserAo)(nil), func() api.UserDb {
		return &UserDbStub{}
	}).(api.UserAo)
	left := userAo.Add(api.User{})
	right := 10002
	if left != right {
		t.Errorf("Error!")
	}
}

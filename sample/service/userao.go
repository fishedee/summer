package service

import (
	"github.com/fishedee/summer/ioc"
	"github.com/fishedee/summer/sample/api"
)

type UserAoImpl struct {
	userDb api.UserDb
}

func (this *UserAoImpl) Get(id int) api.User {
	return this.userDb.Get(id)
}

func (this *UserAoImpl) Add(data api.User) int {
	return this.userDb.Add(data)
}

func NewUserAoImpl(userDb api.UserDb) api.UserAo {
	userAo := &UserAoImpl{}
	userAo.userDb = userDb
	return userAo
}

func init() {
	ioc.Register(NewUserAoImpl)
}

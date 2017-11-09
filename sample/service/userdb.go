package service

import (
	"github.com/fishedee/summer/ioc"
	"github.com/fishedee/summer/sample/api"
)

type UserDbImpl struct {
	totalId int
	data    map[int]api.User
}

func (this *UserDbImpl) Get(id int) api.User {
	return this.data[id]
}

func (this *UserDbImpl) Add(data api.User) int {
	this.totalId++
	data.Id = this.totalId
	this.data[this.totalId] = data
	return this.totalId
}

func NewUserDbImpl() api.UserDb {
	userDb := &UserDbImpl{}
	userDb.totalId = 10000
	userDb.data = map[int]api.User{}
	return userDb
}

func init() {
	ioc.Register(NewUserDbImpl)
}

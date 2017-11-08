package main

import (
	"github.com/fishedee/summer/ioc"
)

type UserAoImpl struct {
	totalId int
	data    map[int]User
}

func (this *UserAoImpl) Get(id int) User {
	return this.data[id]
}

func (this *UserAoImpl) Add(data User) int {
	this.totalId++
	this.data[this.totalId] = data
	return this.totalId
}

func NewUserAoImpl() UserAo {
	userAo := &UserAoImpl{}
	userAo.totalId = 10001
	userAo.data = map[int]User{}
	return userAo
}

func init() {
	ioc.Register(NewUserAoImpl())
}

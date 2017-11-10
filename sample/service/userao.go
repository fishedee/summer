package service

import (
	"github.com/fishedee/summer/ioc"
	"github.com/fishedee/summer/sample/api"
	"github.com/fishedee/summer/sample/util"
)

type userAoImpl struct {
	userDb api.UserDb
}

func (this *userAoImpl) Get(id int) api.User {
	util.MyLog.Debug("Get %v", id)
	return this.userDb.Get(id)
}

func (this *userAoImpl) Add(data api.User) int {
	util.MyLog.Debug("Add %v", data)
	return this.userDb.Add(data)
}

func newUserAoImpl(userDb api.UserDb) api.UserAo {
	userAo := &userAoImpl{}
	userAo.userDb = userDb
	return userAo
}

func init() {
	ioc.Register(newUserAoImpl)
}

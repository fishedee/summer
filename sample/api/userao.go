package api

type UserAo interface {
	Get(id int) User
	Add(data User) int
}

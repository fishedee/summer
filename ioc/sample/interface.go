package main

type User struct {
	Id   int
	Name string
}

type UserAo interface {
	Get(id int) User
	Add(data User) int
}

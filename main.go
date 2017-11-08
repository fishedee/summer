package main

import (
	"fmt"
)

type UserAo interface {
	Get(id int) int
}

func main() {
	fmt.Println(UserAo(nil))
}

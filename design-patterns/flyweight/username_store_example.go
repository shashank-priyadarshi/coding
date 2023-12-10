package main

import (
	"fmt"
	"strings"
)

var Users []string

type User struct {
	names []uint8
}

func (user User) String() string {
	parts := make([]string, len(user.names))
	for i, v := range user.names {
		parts[i] = Users[v]
	}

	return strings.Join(parts, " ")
}

func NewUser(fullName string) User {
	getOrAdd := func(s string) uint8 {
		for i, v := range Users {
			if s == v {
				return uint8(i)
			}
		}
		Users = append(Users, s)
		return uint8(len(Users) - 1)
	}
	parts := strings.Split(fullName, " ")
	user := User{}
	for _, v := range parts {
		user.names = append(user.names, getOrAdd(v))
	}
	return user
}

func main() {
	user := NewUser("sufi Mirza")
	fmt.Printf("user.String(): %v\n", user.String())
}

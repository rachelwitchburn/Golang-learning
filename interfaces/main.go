package main

import "fmt"

func main () {
	u := User {
		name: "Raquel",
		userName: "rachelwitchburn",
		online: true,
	}

	adm := Admin {
		User{
			name: "Luiz",
			userName: "luizjmd",
			online: true,
		},
		22,
	}

	ShowUserInfo(u)
	ShowUserInfo(adm)
}

func ShowUserInfo(u UsersInterface) {
	fmt.Println(u.Show())
}

type UsersInterface interface {
	Show() string
}

type User struct {
	name string
	userName string
	online bool
}

func (u User) Show() string  {
	return fmt.Sprintf(
		"Hello, my name is %s, and my username is %s.",
		u.name, u.userName,
	)
}

type Admin struct {
	User
	age int
}

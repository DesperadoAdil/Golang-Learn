package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending email to <%s>user: %s<%s>\n", a.level, a.name, a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "adil",
			email: "adil@email.com",
		},
		level: "admin",
	}
	ad.user.notify()
	ad.notify()
}

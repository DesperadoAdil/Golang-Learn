package main

import "fmt"

type notifier interface {
    notify()
}

type user struct {
    name string
    email string
}

func (u *user) notify()  {
    fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

type admin struct {
    name string
    email string
}

func (a *admin) notify()  {
    fmt.Printf("Sending admin to %s<%s>\n", a.name, a.email)
}

func main()  {
    u := user{"adil", "adil@email.com"}
    sendNotification(&u)

    a := admin{"admin", "admin@email.com"}
    sendNotification(&a)
}

func sendNotification(n notifier)  {
    n.notify()
}

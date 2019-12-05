package main

import "fmt"

type notifier interface {
    notify()
}

type user struct {
    name string
    email string
}

/*
// cannot use u (type user) as type notifier in argument to sendNotification:
//    user does not implement notifier (notify method has pointer receiver)

func (u *user) notify()  {
    fmt.Printf("*Sending email to %s<%s>\n", u.name, u.email)
}
*/

func (u user) notify()  {
    fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

func main()  {
    u := &user{"adil", "adil@email.com"}
    sendNotification(u)
}

func sendNotification(n notifier)  {
    n.notify()
}

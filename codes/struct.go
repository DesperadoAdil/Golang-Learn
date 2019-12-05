package main

import (
	"fmt"
)

type user struct {
    name string
    email string
}

func main() {
    adil := user{
        name: "adil",
        email: "adil@email.com",
    }
    //var www user
    var www user = user{
        name: "www",
        email: "www@email.com",
    }
    //www.name = "www"
    //www.email = "www@email.com"
    fmt.Printf("%v in %p\n", adil, &adil)
    fmt.Printf("%v in %p\n", www, &www)
}

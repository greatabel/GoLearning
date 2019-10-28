package main

import (
    "fmt"
)

type user struct {
    name  string
    email string
}

type admin struct {
    person user
    level string
}

func (u user) notify() {
    fmt.Printf("Sending User Email To %s<%s>\n",
        u.name,
        u.email)
}

func (u *user) changeEmail(email string) {
    u.email = email
}

func main() {

    fred := admin {
        person: user{
        name: "Lisa#",
        email: "Lisa#@gmail.com",
        },
        level: "super",

    }
    fred.person.notify()
}
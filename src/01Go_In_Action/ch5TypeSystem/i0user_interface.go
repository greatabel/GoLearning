package main

import (
    "fmt"
)

type notifier interface {
    notify()
}

type user struct {
    name  string
    email string
}

func (u user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n",
        u.name,
        u.email)
}

func (u *user) changeEmail(email string) {
    u.email = email
}

func main() {
    bill := user{"Bill", "bill@gmail.com"}
    sendNotification(bill)

    lisa := &user{"Lisa", "lisa@email.com"}
    sendNotification(lisa)
    
    bill.changeEmail("bill@newdomain.com")
    sendNotification(bill)

    // Pointers of type user can be used to call methods
    // declared with a pointer receiver.
    lisa.changeEmail("lisa@newdomain.com")
    sendNotification(lisa)
}

func sendNotification(n notifier) {
    n.notify()
}
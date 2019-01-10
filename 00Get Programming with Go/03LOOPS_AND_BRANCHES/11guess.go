package main

import (
    "fmt"
    "math/rand"
    
)

// Write a guess-the-number program. Make the computer 
// pick random numbers between 1–100 until it guesses your number, 
// which you declare at the top of the program. Display each guess 
// and whether it was too big or too small.

func main() {
    var myguess = 9

    var count = 0
    for {
        
        count++
        var computer_guess = rand.Intn(100)
        fmt.Println("computer guess =", computer_guess)
        if computer_guess == myguess {
            fmt.Println(myguess, " is right! Computer guessed ", count, " times. break")
            break
        } else {
            fmt.Println(computer_guess, " wrong.")
        }
    }

    
}

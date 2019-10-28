// Sample program to show how to show you how to briefly work
// with the sql package.
package main

import (
    "database/sql"

    _ "01Go_In_Action/ch3Package_toolchain/dbdriver/postgres"
)

// main is the entry point for the application.
func main() {
    sql.Open("postgres", "mydb")
}

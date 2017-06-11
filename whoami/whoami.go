package main

import (
        "fmt"
        "os"
)

func main() {
        user := os.Getenv("USER")

        if len(os.Args) > 1 {
                fmt.Fprintf(os.Stderr, "To many arguments, usage: whoami \n")
                os.Exit(1)
        }

        fmt.Println(user)
}

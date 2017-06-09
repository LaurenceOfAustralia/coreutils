package main

import (
        "fmt"
        "os"
)

func main() {
        pwd := os.Getenv("PWD")

        if len(pwd) > 0 && pwd[0] == '/' {
                fmt.Println(pwd)
        }
}

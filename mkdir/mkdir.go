package main

import (
        "fmt"
        "os"
)

// Simple mkdir. No flags, by default works like GNU mkdir with -p flag added.

func main() {
        if len(os.Args) < 2 {
                fmt.Fprintf(os.Stderr, "Not enough arguments given. Needed to recieve at least two arguments. \n")
                os.Exit(1)
        }

        arg1 := string(os.Args[1])

        err := os.MkdirAll(arg1, 0700)

        if err != nil {
                fmt.Fprintf(os.Stderr, "%s \n", err)
                os.Exit(1)
        }
}

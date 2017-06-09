package main

import (
        "fmt"
        "os"
)

func main() {
        for i := range os.Args {

                if i != 0 {
                        fmt.Print(os.Args[i] + " ")

                        if i == len(os.Args)-1 {
                                fmt.Print(" \n")
                        }

                }    

        }
}

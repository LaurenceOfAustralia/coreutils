package main

import (
        "fmt"
        "os"
        "io/ioutil"
)

// A better named simpler version of cat without 50 strange unnecessary flags.

func main() {
        if len(os.Args) < 2 {
                fmt.Fprintf(os.Stderr, "Not enough arguments given. Needed to recieve at least two arguments. \n")
                os.Exit(1)
        }

        for i, file := range os.Args {
                if i != 0 {
                        data, err := ioutil.ReadFile(file)

                        if err != nil {
                                fmt.Fprintf(os.Stderr, "%s \n", err)
                                os.Exit(1)
                        }
                        fmt.Println(string(data))
                }
        }
}

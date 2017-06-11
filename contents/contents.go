package main

import (
        "fmt"
        "os"
        "io/ioutil"
)

// A better named simpler version of cat without 50 strange unnecessary flags.

func main() {
        for i, file := range os.Args {
                if i != 0 {
                        data, err := ioutil.ReadFile(file)

                        if err != nil {
                                fmt.Fprintf(os.Stderr, "%s \n", err)
                        }
                        fmt.Println(string(data))
                }
        }
}

package main

import (
        "fmt"
        "os"
        "io/ioutil"
)

func main() {
        for i, file := range os.Args {
                if i != 0 {
                        fmt.Println(file)
                        data, err := ioutil.ReadFile(file)

                        if err != nil {
                                fmt.Fprintf(os.Stderr, "%s \n", err)
                        }
                        fmt.Println(string(data))
                }
        }
}

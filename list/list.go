package main

import (
        "fmt"
        "io/ioutil"
        "os"
)

// Extrordinarily simple, no arguments similar to ls.
// TODO add formatting to output.

func main() {
        path := os.Args[1]
        entries, err := ioutil.ReadDir(path)

        if err != nil {
                fmt.Fprintf(os.Stderr, "list: "+path+": "+" No such file or directory \n")
                os.Exit(1)
        }

        for _, f := range entries {
                fmt.Println(f.Name())
        }
}

package main

import (
        "fmt"
        "io/ioutil"
        "os"
)

// Extrordinarily simple, no arguments similar to ls.
// TODO add formatting to output.

func main() {

        if len(os.Args) < 2 {
                fmt.Fprintf(os.Stderr, "Not enough arguments given. Needed to recieved at least two arguments. \n")
                os.Exit(1)
        }
        
        path := os.Args[1]
        entries, err := ioutil.ReadDir(path)

        if err != nil {
                fmt.Fprintf(os.Stderr, "list: "+path+": "+" No such file or directory \n")
                os.Exit(1)
        }

        for i, f := range entries {
                // Terrible hacky formatting. Fix sometime.
                fmt.Fprintf(os.Stdout, "%s    ", f.Name())

                if i == len(entries)-1 {
                        fmt.Println("")
                }
        }
}

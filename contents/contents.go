package main

import (
        "fmt"
        "os"
        "io"
        "bufio"
)

func main() {
        if len(os.Args) < 2 {
                fmt.Fprintf(os.Stderr, "Not enough arguments given. Needed to recieve at least two arguments. \n")
                os.Exit(1)
        }

        for i, file := range os.Args {
                if i != 0 {
                        /* This is the most efficiant method i have found so far.
                        However it loads the entire file into memory. This is bad for big files*/
                        data, err := os.Open(file)

                        if err != nil {
                                fmt.Fprintf(os.Stderr, "%s \n", err)
                                os.Exit(1)
                        }

                        defer data.Close()

                        r := bufio.NewReader(data)

                        for {
                                path, err := r.ReadString(10) // 0x0A separator = newline

                                fmt.Println(path)

                                if err == io.EOF {
                                        break
                                        } else if err != nil {
                                                println(err)
                                        }
                                }
                }
        }
}

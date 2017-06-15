package main

import (
        "fmt"
        "io/ioutil"
        "os"
)

// Simple copy of ls which by default shows hidden folders.

func main() {
        path := os.Getenv("PWD")

        if len(os.Args) > 1 {
                path = os.Args[1]
        }

        entries, err := ioutil.ReadDir(path)

        if err != nil {
                fmt.Fprintf(os.Stderr, "list: "+path+": "+" No such file or directory \n")
                os.Exit(1)
        }

        output := make([]string, 1)
        current_slice_num := 0

        for i, f := range entries {
                item := f.Name()

                // The formatting of the output works by the usage of tab charichters. aka \t
                if len(item) + len(output[current_slice_num]) + 1 < 70 {
                        output[current_slice_num] = output[current_slice_num] + item + "\t"
                } else {
                        output[current_slice_num] = output[current_slice_num] + "\t"
                        output = append(output, item + "\t")
                        current_slice_num = current_slice_num + 1
                }

                if i == len(entries)-1 {
                        for i := range output {
                                fmt.Println(output[i])
                        }
                }
        }
}

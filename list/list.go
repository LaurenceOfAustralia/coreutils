package main

import (
        "fmt"
        "io/ioutil"
        "os"
)

// List the flag -h shows hidden files not -a as in ls.

func main() {
        flags := set_flag_values()
        path := os.Getenv("PWD")

        if len(os.Args) > 1 && flags == nil {
                path = os.Args[1]
        }

        if len(os.Args) > 1 && flags != nil {
                path = os.Args[2]
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

                if flags["h"] == false {
                        if string([]rune(item)[0]) == "." {
                                continue
                        }
                }

                // TODO The formatting of the output works by the usage of tab charichters.
                // Not great but could be worse. Work on it.
                if len(item) + len(output[current_slice_num]) + 1 < 90 {
                        output[current_slice_num] = output[current_slice_num] + item + "\t "
                } else {
                        output[current_slice_num] = output[current_slice_num] + "\t "
                        output = append(output, item + "\t ")
                        current_slice_num = current_slice_num + 1
                }

                if i == len(entries)-1 {
                        for i := range output {
                                fmt.Println(output[i])
                        }
                }
        }
}

func set_flag_values()(result map[string]bool) {
        result = make(map[string]bool)

        if len(os.Args) > 1 && os.Args[1][0:1] == "-" && len(os.Args[1]) > 1 {
                println("We have a flag!!")

                for _, c := range os.Args[1][1:len(os.Args[1])] {
                        char := string(c)

                        if char != " " {
                                result[char] = true
                        }
                }
                return result
        }

        return result
}

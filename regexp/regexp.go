package main

import (
        "fmt"
        "os"
        "bufio"
        "regexp"
)
// Regexp matcher returns true if there are matches otherwise false.

func main() {
        if len(os.Args) < 2 {
                fmt.Fprintf(os.Stderr, "Not enough arguments given. Needed to recieve at least two arguments. \n")
                os.Exit(1)
        }

        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()

        input := scanner.Text()
        arg1 := os.Args[1]

        matched, err := regexp.MatchString(arg1, input)
        if err != nil {
                fmt.Fprintf(os.Stderr, "%s"+"\n", err)
                os.Exit(1)
        }

	fmt.Println(matched)

}

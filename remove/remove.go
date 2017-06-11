package main

import (
        "fmt"
        "os"
)

// A new version of rm, by default it deletes directorys.
// TODO add trash in ~/remove/date deleted/deleted item. Also add command to clear trash.

func main() {
        if len(os.Args) < 2 {
                fmt.Fprintf(os.Stderr, "Not enough arguments given. Needed to recieve at least two arguments. \n")
                os.Exit(1)
        }

        arg1 := os.Args[1]

        err := os.Remove(arg1)

        if err != nil {
                fmt.Fprintf(os.Stderr, "%s \n", err)
                os.Exit(1)
        }
}

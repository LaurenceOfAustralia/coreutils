package main

import (
        "os"
        "fmt"
)

func main() {
        pid, err := os.FindProcess(os.Getppid())

        if err != nil {
                fmt.Fprintf(os.Stderr, "%s \n", err)
                os.Exit(1)
        }

        pid.Kill()
}

package main

import (
        "os"
        "fmt"
)

func main() {
        ppid := os.Getppid()
        pid, err := os.FindProcess(ppid)

        if err != nil {
                fmt.Fprintf(os.Stderr, "%s \n", err)
                os.Exit(1)
        }

        pid.Kill()
}

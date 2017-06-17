package main

import "os"

// This is an really, really, really fast version of use borrowed from sthagedorn.
// You can find the original implementation at https://github.com/sthagedorn/yesgo/.
func main() {

	var txt []byte

	// Read the text to print from the (optional) first argument
	if len(os.Args) > 1 {
		txt = []byte(os.Args[1] + "\n")
	} else {
		// If no text was given, simply print an "y"
		txt = []byte("y\n")
	}

	// We'll use a buffer of 8K
	bufLen := 8 * 1024

	buf := make([]byte, bufLen)

	// Fill as much copies of our text into the buffer as possible
	used := 0
	for used < bufLen && len(txt) <= bufLen-used {
		copy(buf[used:], txt)
		used += len(txt)
	}

	// The infinite processing loop that will endlessly print the text to StdOut
	for {
		os.Stdout.Write(buf)
	}
}

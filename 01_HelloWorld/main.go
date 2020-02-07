// executable program
package main

import (
	"fmt"
	f "fmt"
	"runtime"
)

const ok = true

// main function
func main() {
	var hello = "Hello GO!"
	f.Println(hello, ok)
	bye()

	// Print number of CPUs
	fmt.Println("You have:", runtime.NumCPU(), "CPUs")
}

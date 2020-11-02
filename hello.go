package main // like namespace to group functions

import (
	"fmt" // fmt package that has functions for

	"rsc.io/quote" // go.mod file lists the specific modules and versions providing those packages
	//  go mod init hello
)

// `go run` will call main() "by default"
func main() {
	fmt.Println("Hello, World") //formatting text and displaying it to console

	fmt.Println(quote.Go())
}

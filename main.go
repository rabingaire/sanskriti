package main

import (
	"fmt"
	"github.com/rabingaire/sanskriti/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Sanskriti Programming Language build by @rabingaire! \n", user.Username)
	fmt.Printf("To Exit type ^C Thank You! \n")
	repl.Start(os.Stdin, os.Stdout)
}

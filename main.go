package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rabingaire/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Monkey Programming Language\n", user.Username)
	fmt.Printf("To Exit type ^C Thank You! \n")
	repl.Start(os.Stdin, os.Stdout)
}

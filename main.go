package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	user "os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
	"monkey/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free too type in commands")
	repl.Start(os.Stdin, os.Stdout)
}

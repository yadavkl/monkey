package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/yadavkl/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %v this i monkey programming langauges\n", user)
	fmt.Println("Feel free to type commands")
	repl.Start(os.Stdin, os.Stdout)
}

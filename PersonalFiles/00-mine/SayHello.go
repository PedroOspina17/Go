package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Hello world !")
	if len(args) > 1 {
		fmt.Println("Hello ", args[1:])
	} else {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please type your name: ")
		text, _ := reader.ReadString('\n')

		fmt.Println("Hello", text)
	}

}

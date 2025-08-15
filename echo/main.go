package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, i := range args {
		fmt.Printf("%s ", i)
	}

}

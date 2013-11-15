package main

import (
	"fmt"
	"github.com/laher/wget-go/wget"
	"os"
)

func main() {
	err := wget.Wget(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

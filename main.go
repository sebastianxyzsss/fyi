package main

import (
	"fmt"
	"os"

	"github.com/sebastianxyzsss/fyi/fyi"
)

func main() {
	fmt.Println("---------------------------- fyi")

	var filePath string = ""

	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	fyi.RunHttpServer(filePath, "8888")

	fmt.Println("...")
}

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

	fyi.RunHttpServer(filePath, style, "8888")

	fmt.Println("...")
}

var style string = `
.darkmode {
	background-color: black;
	color: white;
}
.darkmode a {
	color: hotpink;
	font-size: 1.1em;
}
code {
	font-family: Consolas,"courier new";
	color: hotpink;
	background: black;
	border-style: dotted;
	border-color: darkgray;
	border-width: 0.01em;
	padding: 0.3em;
	font-size: 1.1em;
}
`

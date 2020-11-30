package main

import (
	"fmt"
	"os"

	"github.com/gobwas/glob"
)

func main() {
	var g glob.Glob

	// Store args
	inputPaths := os.Args[1:]
	fmt.Println(inputPaths)

	for path := range inputPaths {
		g = glob.MustCompile(inputPaths[path])
		fmt.Println(g.Match(inputPaths[path])) // true

	}

	// provide file path with wildcard
	// confirm glob match based off provided path(s)
	// ex /var/opt/*.log
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Nikkely/germaid/internal/graph"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	var stdin string
	for scanner.Scan() {
		stdin += fmt.Sprintln(scanner.Text())
	}

	stdout, err := graph.ParseNaively(stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, stdout)
	os.Exit(0)
}

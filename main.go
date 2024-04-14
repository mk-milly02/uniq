package main

import (
	"flag"
	"fmt"
	"uniq/uniq"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	input := uniq.ReadFromFile(filename)
	output := uniq.Unique(input)

	for _, line := range output {
		fmt.Print(line)
	}
}

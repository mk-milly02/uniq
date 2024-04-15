package main

import (
	"flag"
	"uniq/uniq"
)

func main() {
	c := flag.Bool("c", false, "adds a new first column that contains a count of the number of times a line appears in an input.")
	flag.Parse()

	var filename string
	var output_filename string
	var output []string
	var text []byte

	if filename = flag.Arg(0); filename == "-" {
		text = uniq.ReadFromStandardInput()
	} else {
		text = uniq.ReadFromFile(filename)
	}

	if *c {
		output = uniq.UniqueWithCount(text)
	} else {
		output = uniq.Unique(text)
	}

	if output_filename = flag.Arg(1); output_filename != "" {
		uniq.WriteToFile(output_filename, output)
	} else {
		uniq.WriteToStandardOutput(output)
	}

}

package main

import (
	"flag"
	"slices"
	"uniq/uniq"
)

func main() {
	c := flag.Bool("c", false, "adds a new first column that contains a count of the number of times a line appears in an input.")
	d := flag.Bool("d", false, "outputs only repeated lines")
	u := flag.Bool("u", false, "outputs only unique lines")
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

	if *d {
		output = uniq.UniqueRepeatedLines(text)
	}

	if *u {
		output = uniq.UniqueOnly(text)
	}

	if *c && *d {
		o1 := uniq.UniqueWithCount(text)
		o2 := uniq.UniqueRepeatedLines(text)
		o3 := uniq.UniqueWithCountS(o2)
		output = slices.Concat(o1, o3)
	}

	if output_filename = flag.Arg(1); output_filename != "" {
		uniq.WriteToFile(output_filename, output)
	} else {
		uniq.WriteToStandardOutput(output)
	}

}

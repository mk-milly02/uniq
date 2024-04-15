package main

import (
	"flag"
	"uniq/uniq"
)

func main() {
	flag.Parse()

	var filename string
	var output_filename string
	var text []byte

	if filename = flag.Arg(0); filename == "-" {
		text = uniq.ReadFromStandardInput()
	} else {
		text = uniq.ReadFromFile(filename)
	}

	output := uniq.Unique(text)

	if output_filename = flag.Arg(1); output_filename != "" {
		uniq.WriteToFile(output_filename, output)
	} else {
		uniq.WriteToStandardOutput(output)
	}

}

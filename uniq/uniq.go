package uniq

import (
	"bytes"
	"io"
	"log"
	"os"
	"slices"
)

func ReadFromFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s: no such file or directory", filename)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func Unique(text []byte) (output []string) {
	reader := bytes.NewBuffer(text)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !slices.Contains(output, line) {
			output = append(output, line)
		}
	}
	return output
}

package uniq

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
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

func ReadFromStandardInput() []byte {
	text, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func WriteToStandardOutput(output []string) {
	for _, line := range output {
		fmt.Print(line)
	}
}

func WriteToFile(filename string, output []string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range output {
		_, err := file.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Output has been written to file %s", filename)
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

func UniqueWithCount(text []byte) []string {
	var output []string
	lines_with_count := make(map[string]int)
	reader := bytes.NewBuffer(text)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var lines []string
		for k := range lines_with_count {
			lines = append(lines, k)
		}
		if !slices.Contains(lines, line) {
			lines_with_count[line] = 1
		} else {
			c := lines_with_count[line]
			c++
			lines_with_count[line] = c
		}
	}

	for k, v := range lines_with_count {
		line := strconv.Itoa(v) + " " + k
		output = append(output, line)
	}
	return output
}

// Returns only repeated lines
func UniqueRepeatedLines(text []byte) (output []string) {
	var no_duplicates []string
	reader := bytes.NewBuffer(text)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !slices.Contains(no_duplicates, line) {
			no_duplicates = append(output, line)
		} else {
			output = append(output, line)
		}
	}
	return output
}

// Returns only uniquw lines
func UniqueOnly(text []byte) (output []string) {
	reader := bytes.NewBuffer(text)
	duplicates := UniqueRepeatedLines(text)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !slices.Contains(duplicates, line) {
			output = append(output, line)
		}
	}
	return output
}

package uniq

import (
	"bytes"
	"fmt"
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

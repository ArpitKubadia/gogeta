package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var in io.Reader

	// define command-line flags
	filePtr := flag.String("file", "", "input file path")
	urlPtr := flag.String("url", "", "input URL")
	modePtr := flag.String("mode", "all", "Values: all, subdomain, port")

	flag.Parse()

	// check which flag is provided
	switch {
	case *filePtr != "":
		// open input file
		file, err := os.Open(*filePtr)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		in = file
	case *urlPtr != "":
		// read from input URL
		in = strings.NewReader(*urlPtr)
	default:
		// read from standard input
		in = os.Stdin
	}

	// create scanner for input
	inputs := []string{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	var funcMap = map[string]interface{}{
		"subdomain": subdomainEnum,
		"port":      portScanning,
	}

	results := funcMap[*modePtr].(func([]string) []string)(inputs)
	fmt.Println(results)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

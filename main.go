package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var in io.Reader

	// define command-line flags
	filePtr := flag.String("file", "", "input file path")
	urlPtr := flag.String("url", "", "input URL")
	tagPtr := flag.String("mode", "all", "Values: all, subdomain, port, host, nuclei")
	outputFile := flag.String("o", "", "Path to the output file")

	// tagPtr := flag.String("tag", "", "input URL")

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

	// var funcMap = map[string]interface{}{
	// 	"subdomain": subdomainEnum,
	// 	"port":      portScanning,
	// 	"host":      hostScanning,
	// 	"nuclei":    nucleiScanning,
	// 	"execCmd":   execCommands,
	// }

	// var tags []string
	// if *tagPtr == "all" {
	// 	tags = []string{"subdomain", "port", "hosts", "nuclei"}
	// } else {
	// 	tags = make([]string, 1)
	// 	tags[0] = *tagPtr
	// }
	// fmt.Println(tags)

	var file *os.File
	var err error

	if *outputFile != "" {
		file, err = os.Create(*outputFile)
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
		}
		defer file.Close()
	} else {
		file = os.Stdout
	}

	for _, domain := range inputs {
		// fmt.Println(domain)
		results := execCommands(domain, *tagPtr)
		output := make([]string, 0, len(results))
		for k := range results {
			if k != "" {
				output = append(output, k)
				fmt.Fprintln(file, k)
			}

		}

		// fmt.Println(output)
		// fmt.Println(len(output))
		// results = append(results, subdomain_scan(domain)...) //to resolve "cannot use scan(domain) (value of type []string) as string value in argument to append"
	}
	// results := funcMap[*tagPtr].(func([]string) []string)(inputs)
	// fmt.Println(results)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

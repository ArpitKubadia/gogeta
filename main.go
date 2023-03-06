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
	// allPtr := flag.Bool("all", false, "Scan Everything")
	// subdomainPtr := flag.Bool("subdomain", false, "Scan Subdomains")
	// portPtr := flag.Bool("port", false, "Scan Ports")

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

	switch {
	case *modePtr == "subdomain":
		results := subdomainEnum(inputs)
		fmt.Println(results)
	}

	// if *allPtr {
	// 	// scan input line by line
	// 	for scanner.Scan() {
	// 		// process each line
	// 		line := scanner.Text()
	// 		fmt.Println(line)
	// 		subdomains := subdomainEnum(line)
	// 		// fmt.Println(subdomains)
	// 		for i, subdomain := range subdomains {
	// 			fmt.Println(i, subdomain)
	// 			ports := portScanning(subdomain)
	// 			// fmt.Println(ports)
	// 			for j, port := range ports {
	// 				fmt.Println(j, port)
	// 			}
	// 		}

	// 	}

	// }

	// if *subdomainPtr {
	// 	// scan input line by line
	// 	for scanner.Scan() {
	// 		// process each line
	// 		line := scanner.Text()
	// 		fmt.Println(line)
	// 		subdomains := subdomainEnum(line)
	// 		fmt.Println(subdomains)
	// 	}
	// }

	// if *portPtr {

	// }

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func port_scan(subdomain string) []string {
	cmd := exec.Command("naabu", "-host", subdomain)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(out.String(), "\n")
}

func portScanning(input []string) []string {
	results := []string{}
	for _, subdomain := range input {
		fmt.Println(subdomain)
		results = append(results, port_scan(subdomain)...) //to resolve "cannot use scan(domain) (value of type []string) as string value in argument to append"
	}
	return results
}

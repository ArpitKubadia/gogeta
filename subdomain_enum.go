package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func subdomain_scan(domain string) []string {
	cmd := exec.Command("assetfinder", domain)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(out.String(), "\n")
}

func subdomainEnum(input []string) []string {
	results := []string{}
	for _, domain := range input {
		fmt.Println(domain)
		results = append(results, subdomain_scan(domain)...) //to resolve "cannot use scan(domain) (value of type []string) as string value in argument to append"
	}
	return results
}

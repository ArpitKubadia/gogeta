package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func subdomainEnum(domain string) []string {
	cmd := exec.Command("assetfinder", domain)
	fmt.Println(cmd.String())
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(out.String(), "\n")
}

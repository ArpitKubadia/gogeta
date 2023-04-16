package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func nuclei_scan(subdomain string) []string {
	cmd := exec.Command("nuclei", "-u", subdomain, "-silent", "-c", "200", "-rl", "1000", "-nc")
	fmt.Println(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start the command asynchronously.
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Use a scanner to read the command's output in real time.
	scanner := bufio.NewScanner(stdout)
	var out strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		out.WriteString(line)
		out.WriteString("\n")
		fmt.Println(line)
	}

	// Wait for the command to finish and check for errors.
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return strings.Split(out.String(), "\n")
}

func nucleiScanning(input []string) []string {
	results := []string{}
	for _, subdomain := range input {
		fmt.Println(subdomain)
		results = append(results, nuclei_scan(subdomain)...) //to resolve "cannot use scan(domain) (value of type []string) as string value in argument to append"
	}
	return results
}

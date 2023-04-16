package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Command struct {
	Cmd string `yaml:"cmd"`
}

type Config struct {
	Tag      string    `yaml:"tag"`
	Commands []Command `yaml:"commands"`
}

func execCommands(domain string, tag_input string) []string {
	// Get the tag from command line arguments
	if len(os.Args) < 2 {
		log.Fatal("Please provide a tag as an argument")
	}

	// Variables to be passed dynamically to the Go program
	vars := map[string]string{
		"domain": domain,
	}
	var tags []string
	if tag_input == "all" {
		// Collect all unique tags from YAML files
		folder := "."
		uniqueTags := make(map[string]struct{})
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
				config, err := readYAMLFile(path)
				if err != nil {
					return err
				}
				uniqueTags[config.Tag] = struct{}{}
			}
			return nil
		})

		if err != nil {
			log.Fatalf("Error collecting tags: %v", err)
		}

		for tag := range uniqueTags {
			tags = append(tags, tag)
		}
	} else {
		// Use the provided tags
		tags = make([]string, 1)
		tags[0] = tag_input
	}

	fmt.Println(tags)

	for _, tag := range tags {
		fmt.Printf("Processing tag: %s\n", tag)

		// Read YAML files from the folder
		folder := "." // Set the folder containing YAML files
		outputs := []string{}
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
				config, err := readYAMLFile(path)
				if err != nil {
					return err
				}

				if config.Tag == string(tag) {
					output, err := executeCommands(config, vars)
					if err != nil {
						return err
					}
					outputs = append(outputs, output...)
				}
			}
			return nil
		})

		if err != nil {
			log.Fatalf("Error processing YAML files: %v", err)
		}

		// Print the list of outputs for the current tag
		fmt.Printf("Outputs for tag %s:\n", tag)
		for _, output := range outputs {
			fmt.Printf(output)
		}
		fmt.Println()
	}
	results := []string{}
	return results
}

func readYAMLFile(filepath string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %v", err)
	}

	return &config, nil
}

func renderTemplate(tmpl string, vars map[string]string) (string, error) {
	t, err := template.New("command").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var rendered strings.Builder
	err = t.Execute(&rendered, vars)
	if err != nil {
		return "", err
	}

	return rendered.String(), nil
}

func executeCommands(config *Config, vars map[string]string) ([]string, error) {
	outputs := []string{}
	for _, command := range config.Commands {
		cmdStr, err := renderTemplate(command.Cmd, vars)
		if err != nil {
			return nil, fmt.Errorf("Error rendering command template: %v", err)
		}

		output, err := executeCommand(cmdStr)
		if err != nil {
			return nil, fmt.Errorf("Error executing command: %v", err)
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func executeCommand(cmdStr string) (string, error) {
	cmd := exec.Command("sh", "-c", cmdStr)
	fmt.Println(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Error executing command: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

#!/bin/bash

# Install naabu
go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest

# Install assetfinder
go install github.com/tomnomnom/assetfinder@latest

# Install subfinder
go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest

# Install httpx
go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest

# Install httprobe
go install github.com/tomnomnom/httprobe@latest

# Install nuclei
go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest

# Install crobat
go install github.com/cgboal/sonarsearch/cmd/crobat@latest

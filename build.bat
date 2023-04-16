@echo off

rem Install naabu
go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest

rem Install assetfinder
go install github.com/tomnomnom/assetfinder@latest

rem Install subfinder
go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest

rem Install httpx
go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest

rem Install httprobe
go install github.com/tomnomnom/httprobe@latest

rem Install nuclei
go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest

rem Install crobat
go install github.com/cgboal/sonarsearch/cmd/crobat@latest

rem Create tools folder if it doesn't exist
if not exist tools mkdir tools

rem Install github-search
cd tools

if not exist github-search (
    git clone https://github.com/gwen001/github-search
    cd github-search
    pip3 install -r requirements.txt
    cd ..
)
cd ..

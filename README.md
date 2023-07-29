# file-utils

file-utils is a Command-Line Interface (CLI) tool developed in Golang to perform various file operations efficiently. This repository contains the source code for the file-utils CLI tool, which allows users to create, read, copy, and delete files from the terminal.

## Features

- **file-create:** Create a new file with the specified name.
- **file-read:** Read and display the content of a file to the console.
- **file-copy:** Copy a file from a source location to a destination location.
- **file-delete:** Delete a file from the filesystem permanently.

## How to Build:
To build the file-utils CLI tool, navigate to the directory containing the main.go file and run the following command in the terminal:

```javascript
go build -o file-utils
```

## Usage/Examples

```javascript
# Create a new file
./file-utils file-create -name demo.go

# Read the content of a file
./file-utils file-read -name demo.go

# Copy a file to another location
./file-utils file-copy -src demo.go -dest main.go

# Delete a file
./file-utils file-delete -name main.go

```


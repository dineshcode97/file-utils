package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Define flags for CLI commands
	createCmd := flag.NewFlagSet("file-create", flag.ExitOnError)
	readCmd := flag.NewFlagSet("file-read", flag.ExitOnError)
	copyCmd := flag.NewFlagSet("file-copy", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("file-delete", flag.ExitOnError)

	// Define flags and their respective variables for each command
	createFileName := createCmd.String("name", "", "Name of the file to create")
	readFileName := readCmd.String("name", "", "Name of the file to read")
	copySrcFileName := copyCmd.String("src", "", "Source file name")
	copyDestFileName := copyCmd.String("dest", "", "Destination file name")
	deleteFileName := deleteCmd.String("name", "", "Name of the file to delete")

	// Check if the first argument is empty or an invalid command
	if len(os.Args) < 2 {
		fmt.Println("Usage: file-utils <command> [<args>]")
		fmt.Println("Available commands: file-create, file-read, file-copy, file-delete")
		os.Exit(1)
	}

	// Parse the command and execute the respective operation
	switch os.Args[1] {
	case "file-create":
		createCmd.Parse(os.Args[2:])
		if *createFileName == "" {
			fmt.Println("Please provide the name of the file to create using the -name flag")
			createCmd.PrintDefaults()
			os.Exit(1)
		}
		err := createFile(*createFileName)
		if err != nil {
			fmt.Println("Error creating the file:", err)
			os.Exit(1)
		}
		fmt.Println("File created successfully:", *createFileName)

	case "file-read":
		readCmd.Parse(os.Args[2:])
		if *readFileName == "" {
			fmt.Println("Please provide the name of the file to read using the -name flag")
			readCmd.PrintDefaults()
			os.Exit(1)
		}
		err := readFile(*readFileName)
		if err != nil {
			fmt.Println("Error reading the file:", err)
			os.Exit(1)
		}

	case "file-copy":
		copyCmd.Parse(os.Args[2:])
		if *copySrcFileName == "" || *copyDestFileName == "" {
			fmt.Println("Please provide both source and destination file names using -src and -dest flags")
			copyCmd.PrintDefaults()
			os.Exit(1)
		}
		err := copyFile(*copySrcFileName, *copyDestFileName)
		if err != nil {
			fmt.Println("Error copying the file:", err)
			os.Exit(1)
		}
		fmt.Println("File copied successfully from", *copySrcFileName, "to", *copyDestFileName)

	case "file-delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteFileName == "" {
			fmt.Println("Please provide the name of the file to delete using the -name flag")
			deleteCmd.PrintDefaults()
			os.Exit(1)
		}
		err := deleteFile(*deleteFileName)
		if err != nil {
			fmt.Println("Error deleting the file:", err)
			os.Exit(1)
		}
		fmt.Println("File deleted successfully:", *deleteFileName)

	default:
		fmt.Println("Invalid command. Available commands: file-create, file-read, file-copy, file-delete")
		os.Exit(1)
	}
}

func createFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func readFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(srcFileName, destFileName string) error {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

func deleteFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}


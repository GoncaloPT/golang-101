package main

import (
	"fmt"
	"log"
	"os"
)

func printDirContents(dirPath string) error {
	dir, err := os.Open("./error-handling")
	if err != nil {
		return fmt.Errorf("error opening dir: %v", err)
	}
	entries, err := dir.ReadDir(0)
	if err != nil {
		return fmt.Errorf("error opening dir: %v", err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	return nil
}

func killServer(pidFilePath string) error {

	file, err := os.Open(pidFilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}(file)

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("error scanning pid: %v", err)
	}

	fmt.Printf("Killing server with pid = %d\n", pid)
	return nil
}

func main() {
	if err := printDirContents("error-handling"); err != nil {
		log.Fatalf("Error printing dir contents: %v", err)
	}
	if err := killServer("error-handling/server.pid"); err != nil {
		log.Fatalf("Error killing server: %v", err)
	}
	os.Exit(1)
}

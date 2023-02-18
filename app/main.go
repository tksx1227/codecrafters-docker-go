package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: run <image> <command> <arg1> <arg2> ...")
		os.Exit(1)
	}
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	if 0 < stdout.Len() {
		fmt.Fprint(os.Stdout, string(stdout.Bytes()))
	}
	if 0 < stderr.Len() {
		fmt.Fprint(os.Stderr, string(stderr.Bytes()))
	}
}

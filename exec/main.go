package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	PATH := os.Getenv("PATH")
	log.Printf("PATH: %v", PATH)
	path, err := exec.LookPath("dot")
	if err != nil {
		log.Fatalf("Look path err: %v", err)
	}
	log.Printf("path: %v\n", path)
	cmd := exec.Command(path)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

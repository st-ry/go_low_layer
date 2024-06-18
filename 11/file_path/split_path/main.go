package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
	fragments := strings.Split(dir, string(filepath.Separator))
	fmt.Printf("array: %v\n", fragments)
}
package main

import (
	"os"
	"path/filepath"
)

func main() {
	//read test data
	var testData = []byte("012345678910ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	err := os.WriteFile(testPath, testData, 0644)
	if err != nil {
		panic(err)
	}
	//map memory
}
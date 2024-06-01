package main

import (
	"crypto/rand"
	"io"
	"os"
	"strings"
)


func main(){
	c := 1024
	b := make([]byte, c)
	
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	
	str := string(b)
	
	reader := strings.NewReader((str))

	file, err := os.Create("rand_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, reader)
	
	//smaple
	file2, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	io.CopyN(file2, rand.Reader, 1024)
	
}

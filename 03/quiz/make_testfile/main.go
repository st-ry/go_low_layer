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
}
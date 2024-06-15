package main

import (
	"fmt"
	"io"
	"os"
)

//create
func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

//read
func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Appended content\n")
}

func makeDir() {
	os.Mkdir("setting", 0644)
	os.MkdirAll("setting/myapp/networksettings", 0644)
}

func main(){
	open()
	read()
	append()
	makeDir()
}
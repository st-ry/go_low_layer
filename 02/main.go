package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main(){
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	//file
	file.Write([]byte("os.File example¥n"))
	file.Close()
	
	//console(std out)
	os.Stdout.Write([]byte("os.Stdout example\n"))
	
	//buffer
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
	
	//net
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	
	conn.Write([]byte("GET / HTTP/1.0\r˜Host:ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}
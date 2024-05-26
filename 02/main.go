package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
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
	
	//bufio
	bufferIo := bufio.NewWriter(os.Stdout)
	bufferIo.WriteString("bufio.Writer ")
	bufferIo.Flush()
	bufferIo.WriteString("example\n")
	bufferIo.Flush()
	
	//net
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r˜Host:ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
	
	//gzip
	gzfile, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(gzfile)
	writer.Header.Name = "text2.txt"
	writer.Write([]byte("gzip.Writer example\n"))
	writer.Close()
	
	//http server
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http.ResponseWriter Sample"))
}
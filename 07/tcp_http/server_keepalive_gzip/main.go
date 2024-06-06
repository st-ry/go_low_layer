package main

import (
	"net/http"
	"strings"
)

//make sure whether client is acceptable zip
func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

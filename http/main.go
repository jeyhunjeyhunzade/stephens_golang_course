package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Our custom type that impelements Writer interface
type logWriter struct {}

func main() {
	resp, err := http.Get("http://google.com")
	if (err != nil) {
		fmt.Println("Error:", err) 
		os.Exit(1)
	}

	// Using Reader interface
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs) // read response (the html content)
	// fmt.Println(string(bs))

	// os.Stdout has Writer interface, resp.Body has Reader interface 
	// io.Copy(os.Stdout, resp.Body) // same thing 


	lw := logWriter{}
	io.Copy(lw, resp.Body)
}


// Writer of logWriter
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just write this many bytes:", len(bs))
	return len(bs), nil
}


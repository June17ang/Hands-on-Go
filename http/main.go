package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	//io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %v bytes", len(bs))
	return len(bs), nil
}
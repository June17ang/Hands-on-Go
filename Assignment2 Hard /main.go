package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct{}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fw := fileWriter{}
	io.Copy(fw, file)
}

func (f fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

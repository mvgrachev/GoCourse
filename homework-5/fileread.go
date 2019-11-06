package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	file, err := os.Open("fileread.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}

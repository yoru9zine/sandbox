package main

import (
	"fmt"
	"log"
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	f, err := os.Open("file")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			pp.Println(err)
			break
		}
		fmt.Printf("%s\n", buf[0:n])
	}
}

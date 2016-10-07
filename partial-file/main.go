package main

import (
	"log"
	"os"
)

func createSparse() {
	f, err := os.Create("./sparse_file")
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	f.WriteAt([]byte{'a'}, 1024*1024)
	f.Close()
}

func createNormal() {
	f, err := os.Create("./data_b")
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	for i := 0; i < 1024*1024; i++ {
		f.WriteAt([]byte{0}, int64(i))
	}
	f.WriteAt([]byte{'a'}, 1024*1024)
	f.Close()
}

func main() {
	createSparse()
	createNormal()
}

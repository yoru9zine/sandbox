package main

import (
	"os"
	"sync"
)

func read(f *os.File, c chan []byte) {
	readbuf := make([]byte, 128)
	for {
		n, err := f.Read(readbuf)
		if err != nil {
			close(c)
			return
		}
		c <- readbuf[:n]
	}
}

func write(c chan []byte, f *os.File) {
	for data := range c {
		f.Write(data)
	}
	f.Close()
}

func main() {
	c := make(chan []byte)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		f, _ := os.Open("./race.go")
		read(f, c)
		wg.Done()
	}()
	go func() {
		f, _ := os.Create("/dev/null")
		write(c, f)
		wg.Done()
	}()
	wg.Wait()
}

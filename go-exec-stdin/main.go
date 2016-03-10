package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	wc, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("stdin pipe: %s", err)
	}
	rc, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("stdout pipe: %s", err)
	}
	cmd.Start()
	wc.Write([]byte("123"))
	wc.Close()
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatalf("read: %s", err)
	}
	println(string(b))
	cmd.Wait()
}

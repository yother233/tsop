package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var (
// arg1 = flag.String("h", "help", "usage")
)

func main() {
	flag.Parse()
	s := flag.Args()
	fmt.Printf("s: %v\n", s)
	// fmt.Printf("arg1: %v\n", *arg1)
	c := exec.Command("git", "clone", s[0])
	if c.Err != nil {
		panic(c.Err)
	}
	fmt.Printf("c.Path: %v\n", c.Path)
	b, err := c.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %v\n", string(b))
}

func clone() {

}

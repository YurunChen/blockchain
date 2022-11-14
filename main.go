package main

//import "fmt"

func main() {
	blockChain := NewBlockChain("班长")
	cli := CLI{blockChain}
	cli.Run()
}

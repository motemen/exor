package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: exor <command> <args>...")
	os.Exit(254)
}

func die(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(255)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	do(os.Args[1], os.Args[2:])
}

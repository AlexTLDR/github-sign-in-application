package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AlexTLDR/github-sign-in-application/hello"
	"github.com/AlexTLDR/github-sign-in-application/world"
)

func main() {
	displayGreetings(os.Stdout)
}

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
}

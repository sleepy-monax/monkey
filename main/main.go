package main

import (
	"monkey/interactive"
	"os"
)

func main() {
	interactive.Start(os.Stdin, os.Stdout)
}

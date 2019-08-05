package main

import (
	"interactive"
	"os"
)

func main() {
	interactive.InteractiveStart(os.Stdin, os.Stdout)
}

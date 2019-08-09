package main

import (
	"fmt"
	"io/ioutil"
	"monkey/interactive"
	"monkey/parser"
	"monkey/token"
	"monkey/tokenizer"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		data, err := ioutil.ReadFile(os.Args[2])

		if err == nil {
			if os.Args[1] == "-t" {
				tok := tokenizer.New(string(data))

				for t := tok.NextToken(); t.Type != token.EOF; t = tok.NextToken() {
					fmt.Printf("%+v\n", t)
				}
			} else if os.Args[1] == "-p" {
				tok := tokenizer.New(string(data))

				pars := parser.New(tok)
				prog := pars.Parse()

				fmt.Printf("ASTDUMP:%+v\n", prog)
			} else if os.Args[1] == "-c" {

			}
		} else {
			fmt.Printf("Error: %s", err.Error())
		}
	}

	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])

		if err == nil {
			tok := tokenizer.New(string(data))

			pars := parser.New(tok)
			prog := pars.Parse()

			fmt.Printf("ASTDUMP:%+v\n", prog)
		}
	} else {
		interactive.Start(os.Stdin, os.Stdout)
	}
}

package interactive

import (
	"bufio"
	"fmt"
	"io"
	"monkey/parser"
	"monkey/tokenizer"
)

const InteractivePrompt = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(InteractivePrompt)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		tok := tokenizer.New(line)
		pars := parser.New(tok)
		prog := pars.Parse()

		errors := pars.Errors
		if len(errors) != 0 {
			fmt.Printf("\033[31mParser has %d errors\033[0m\n", len(errors))

			for _, msg := range errors {
				fmt.Printf("- %s\n", msg)
			}

		}

		fmt.Printf("%+s\n", prog.String())
	}
}

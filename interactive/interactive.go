package interactive

import (
	"bufio"
	"fmt"
	"io"
	"monkey/token"
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
		state := tokenizer.New(line)

		for tok := state.NextToken(); tok.Type != token.EOF; tok = state.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

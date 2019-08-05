package interactive

import (
	"bufio"
	"tokenizer"
	"fmt"
	"io"
)

// InteractivePrompt is the command line promt display in the user terminal.
const InteractivePrompt = ">>"

// InteractiveStart start a new interactive session
func InteractiveStart(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(InteractivePrompt)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		tokenizer := tokenizer.NewTokenizerState(line)

		for token := tokenizer.NextToken(); token.Type != tokenizer.TOKEN_END_OF_FILE; token = tokenizer.NextToken() {
			fmt.Printf("%+v\n", token)
		}
	}
}

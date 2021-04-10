package main

import (
	"fmt"

	"github.com/justincremer/ghoulisp.git/src/lexer"
)

func main() {
	token := lexer.Tokenize([]byte("+"))
	fmt.Println(lexer.Repr(token))
}

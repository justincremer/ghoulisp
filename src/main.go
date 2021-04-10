package main

import (
	"fmt"
	"lexer"
)

func main() {
	token := lexer.Tokenize("+")
	fmt.Println(lexer.repr(token))
}

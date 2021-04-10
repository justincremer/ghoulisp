package lexer

import (
	"fmt"
	"regexp"
)

type Token struct {
	t   string // Type
	val []byte // Value
}

func (token Token) repr() string {
	return fmt.Sprintf("Token: { %v, %v }", token.t, token.val)
}

// Tokenizes a single character, first checking for alphanumeric typing and then for sybolic typing
func Tokenize(c []byte) Token {
	if isNum, _ := regexp.Match("/[0-9]/", c); isNum {
		return Token{t: "num", val: c}
	}
	if isChar, _ := regexp.Match("/[a-z]/i", c); isChar {
		return Token{t: "char", val: c}
	}

	switch string(c) {
	case "(":
		return Token{t: "rprn", val: c}
	case ")":
		return Token{t: "lprn", val: c}
	case "[":
		return Token{t: "lsb", val: c}
	case "]":
		return Token{t: "rsb", val: c}
	case "{":
		return Token{t: "lcb", val: c}
	case "}":
		return Token{t: "rcb", val: c}
	case ".":
		return Token{t: "dot", val: c}
	case ",":
		return Token{t: "comma", val: c}
	case "+":
		return Token{t: "add", val: c}
	case "-":
		return Token{t: "sub", val: c}
	case "/":
		return Token{t: "div", val: c}
	case "*":
		return Token{t: "mul", val: c}
	case "=":
		return Token{t: "eq", val: c}
	case "!":
		return Token{t: "not", val: c}
	case "'":
		return Token{t: "squote", val: c}
	case "\"":
		return Token{t: "dquote", val: c}
	case "\t":
		return Token{t: "tab", val: c}
	case "\n":
		return Token{t: "newline", val: c}
	case " ":
		return Token{t: "whitespace", val: c}
	default:
		return Token{t: "undef", val: c}
	}
}

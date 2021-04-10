package lexer

import (
	"fmt"
	"regexp"
)

type Token struct {
	t   string // Type
	val string // Value
}

// Returns a formated string literal from a token (e.g. "Token: { add, + }")
func (token Token) Repr() string {
	return fmt.Sprintf("Token: { %v, %v }", token.t, token.val)
}

// Tokenizes a single character, first checking for alphanumeric typing and then for sybolic typing
func Tokenize(c string) Token {
	if isNum, _ := regexp.Match("/[0-9]/", []byte(c)); isNum {
		return Token{t: "num", val: c}
	}
	if isChar, _ := regexp.Match("/[a-z]/i", []byte(c)); isChar {
		return Token{t: "char", val: c}
	}

	switch c {
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

package main

import (
	"lang/lex"
	"lang/parser"
	"os"
)

func main() {
	a, _ := os.ReadFile(os.Args[1])
	parser.Parse(lex.Lex(string(a)))

}

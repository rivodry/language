package main

import (
	"encoding/json"
	"fmt"
	"lang/lex"
	"lang/parser"
	"os"
)

func main() {
	a, _ := os.ReadFile(os.Args[1])

	d, _ := json.Marshal(parser.Parse(lex.Lex(string(a))))
	fmt.Println(string(d))

}

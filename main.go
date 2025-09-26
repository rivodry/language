package main

import (
	"fmt"
	"lang/lex"
	"os"
)

func main() {
	a, _ := os.ReadFile(os.Args[1])
	fmt.Println(lex.Lex(string(a)))

}

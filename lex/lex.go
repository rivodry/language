package lex

import (
	"unicode"
)

type Token struct {
	Typ   string
	Value string
}

func ParseString(num *int, Set string) string {
	i := *num
	var str string
	for i < len(Set) && (unicode.IsLetter(rune(Set[i])) || unicode.IsDigit(rune(Set[i])) || Set[i] == '_') {
		str = str + string(Set[i])
		i++
	}
	*num = i
	return str
}

func Lex(Set string) []Token {
	var Tokens []Token
	for i := 0; i < len(Set); {
		if unicode.IsSpace(rune(Set[i])) {
			i++
			continue
		} else if unicode.IsLetter(rune(Set[i])) || Set[i] == '_' {
			str := ParseString(&i, Set)
			Tokens = append(Tokens, Token{"IDENT", str})
		} else {
			i++
		}
	}
	Tokens = append(Tokens, Token{"EOF", ""})
	return Tokens
}

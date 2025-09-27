package lex

import (
	"unicode"
)

type Token struct {
	Typ   string
	Value string
}

func ParseString(num *int, Set string) string {

}
func ParseIdent(num *int, Set string) string {
	i := *num
	var str string
	for i < len(Set) && (unicode.IsLetter(rune(Set[i])) || unicode.IsDigit(rune(Set[i])) || Set[i] == '_') {
		str = str + string(Set[i])
		i++
	}
	*num = i
	return str
}
func ParseNumber(num *int, Set string) string {
	i := *num
	var str string
	var decimal bool
	for i < len(Set) && (unicode.IsDigit(rune(Set[i])) || Set[i] == '.') {
		if Set[i] == '.' && decimal == false {
			decimal = true
			str = str + string(Set[i])
			i++
		} else {
			str = str + string(Set[i])

			i++
		}

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
			str := ParseIdent(&i, Set)
			Tokens = append(Tokens, Token{"IDENT", str})
		} else if unicode.IsDigit(rune(Set[i])) {
			num := ParseNumber(&i, Set)
			Tokens = append(Tokens, Token{"NUMBER", num})
		} else if rune(Set[i]) == '"' {
			str := ParseString(&i, Set)
		} else {
			i++
		}
	}
	Tokens = append(Tokens, Token{"EOF", ""})
	return Tokens
}

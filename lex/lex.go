package lex

import (
	"slices"
	"unicode"
)

type Token struct {
	Typ   string
	Value string
}

func ParseString(num *int, Set string) string {
	i := *num
	i++
	var str string
	for i < len(Set) && (Set[i] != '"') {
		str = str + string(Set[i])
		i++
	}
	i++
	*num = i
	return str
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
func IdentOrKeyWord(str string) Token {
	switch str {
	case "func":
		{
			return Token{"FUNC", ""}
		}
	case "if":
		{
			return Token{"IF", ""}
		}
	case "while":
		{
			return Token{"WHILE", ""}
		}
	case "start":
		{
			return Token{"START", ""}
		}
	case "end":
		{
			return Token{"END", ""}
		}
	case "then":
		{
			return Token{"THEN", ""}
		}
	case "do":
		{
			return Token{"DO", ""}
		}
	case "and":
		{
			return Token{"AND", ""}
		}
	case "or":
		{
			return Token{"OR", ""}
		}

	default:
		{
			return Token{"IDENT", str}
		}
	}
}
func Lex(Set string) []Token {
	var Tokens []Token
	for i := 0; i < len(Set); {
		if unicode.IsSpace(rune(Set[i])) {
			i++
			continue
		} else if unicode.IsLetter(rune(Set[i])) || Set[i] == '_' {
			str := ParseIdent(&i, Set)
			Tokens = append(Tokens, IdentOrKeyWord(str))
		} else if unicode.IsDigit(rune(Set[i])) {
			num := ParseNumber(&i, Set)
			Tokens = append(Tokens, Token{"NUMBER", num})
		} else if rune(Set[i]) == '"' {
			str := ParseString(&i, Set)
			Tokens = append(Tokens, Token{"STRING", str})

		} else if slices.Contains([]rune{'(', ')'}, rune(Set[i])) {

			Tokens = append(Tokens, Token{string(Set[i]), ""})
			i++

		} else if slices.Contains([]rune{'+', '-', '/', '*'}, rune(Set[i])) {
			Tokens = append(Tokens, Token{string(Set[i]), ""})
		} else if Set[i] == '=' {

		} else {
			i++
		}
	}
	Tokens = append(Tokens, Token{"EOF", ""})
	return Tokens
}

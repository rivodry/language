package main

import (
	"encoding/json"
	"fmt"
	"lang/lex"
	"lang/parser"
	"os"
	"strings"
)

func printJSONTree(data interface{}, indent int) {
	prefix := strings.Repeat("  ", indent)

	switch v := data.(type) {
	case map[string]interface{}:

		for key, value := range v {
			fmt.Printf("%s- %s:\n", prefix, key)
			printJSONTree(value, indent+1)
		}
	case []interface{}:

		for i, item := range v {
			fmt.Printf("%s- [%d]:\n", prefix, i)
			printJSONTree(item, indent+1)
		}
	default:

		fmt.Printf("%s  %v\n", prefix, v)
	}
}

func main() {
	a, _ := os.ReadFile(os.Args[1])
	//fmt.Println(lex.Lex(string(a)))
	d, _ := json.Marshal(parser.Parse(lex.Lex(string(a))))
	var jsonData interface{}
	err := json.Unmarshal([]byte(d), &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	printJSONTree(jsonData, 0)

}

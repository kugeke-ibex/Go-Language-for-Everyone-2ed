package main

import (
	"fmt"
	"strings"
)

func Words(s string) string {
	c := len(strings.Fields(s))

	switch {
		case c == 0:
			return "wordless?"
		case c == 1:
			return "one word"
		case c < 4:
			return "a few words"
		case c < 8:
			return "many words"
		default:
			return "too many words"
	}

}

func main() {
	fmt.Println(Words(""))
	fmt.Println(Words("hello"))
	fmt.Println(Words("hello world"))
	fmt.Println(Words("hello world foo bar"))
	fmt.Println(Words("hello world foo bar baz"))
	fmt.Println(Words("hello world foo bar baz qux"))
}

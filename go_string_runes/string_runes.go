package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "congnc"
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("-----------------------------------------")
	fmt.Println("Rune count", utf8.RuneCountInString(s))

	for idx, runValue := range s {
		fmt.Printf("%#U starts at %d\n", runValue, idx)
	}

	fmt.Println("\n Using  DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U start at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'c' {
		fmt.Println("found so sua")
	}
}

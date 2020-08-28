package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]int)
	rm := make(map[string]int)
	fmt.Println("Please enter symbols")
	var enteredSymbols string
	_, err := fmt.Scanf("%s", &enteredSymbols)
	if err != nil {
		panic(err.Error())
	}
	symbolsRune := []rune(enteredSymbols)
	for _, r := range symbolsRune {
		m[strconv.QuoteRune(r)]++
	}
	for s, rp := range m {
		if rp > 1 {
			rm[s] = rp
		}
	}
	if len(rm) == 0 {
		fmt.Println("The string symbols are uniq")
	} else {
		fmt.Println(rm)
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	symbolMap := make(map[string]int)
	reapetSymbolMap := make(map[string]int)
	fmt.Println("Please enter symbols")
	var enteredSymbols string
	_, err := fmt.Scanf("%s", &enteredSymbols)
	if err != nil {
		panic(err.Error())
	}
	symbolsRune := []rune(enteredSymbols)
	for _, r := range symbolsRune {
		symbolMap[strconv.QuoteRune(r)]++
	}
	for s, rp := range symbolMap {
		if rp > 1 {
			reapetSymbolMap[s] = rp
		}
	}
	if len(reapetSymbolMap) == 0 {
		fmt.Println("The string symbols are uniq!")
	} else {
		fmt.Println(reapetSymbolMap)
	}
}

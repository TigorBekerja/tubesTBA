package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Masukkan kalimat: ")
	var kalimat string
	fmt.Scan(&kalimat)

	
	if bacaPredikat(kalimat) {
		fmt.Println("Kalimat termasuk ke predikat")
	} else {
		fmt.Println("Kalimat bukan termasuk ke predikat")
	}
}

func bacaPredikat(kalimat string) bool {
	words := strings.Fields(kalimat)
	for _, word := range words {
		if isPredikat(word) {
			return true
		}
	}
	return false
}

func isPredikat(word string) bool {
	state := 0
	for _, char := range word {
		fmt.Printf("State: %d, Char: %c\n", state, char)
		if state == 0 && char == 'm' {
			state = 1
		} else if state == 1 && char == 'e' {
			state = 2
		} else if state == 2 && char == 'm' {
			state = 3
		} else if state == 2 && char == 'n' {
			state = 4
		} else if state == 2 && char == 'l' {
			state = 7
		} else if state == 3 && char == 'o' {
			state = 5
		} else if state == 3 && char == 'b' {
			state = 6
		} else if state == 3 && char == 'l' {
			state = 7
		} else if state == 4 && char == 'g' {
			state = 8
		} else if state == 5 && char == 't' {
			state = 10
		} else if state == 6 && char == 'u' {
			state = 11
		} else if state == 7 && char == 'i' {
			state = 12
		}  else if state == 8 && char == 'i' {
			state = 13
		} else if state == 8 && char == 'a' {
			state = 9
		} else if state == 9 && char == 'm' {
			state = 14
		} else if state == 10 && char == 'o' {
			state = 15
		} else if state == 11 && char == 'a' {
			state = 16
		} else if state == 12 && char == 'h' {
			state = 11
		} else if state == 13 && char == 'k' {
			state = 17
		} else if state == 14 && char == 'b' {
			state = 18
		} else if state == 15 && char == 'n' {
			state = 19
		} else if state == 16 && char == 't' {
			state = 20
		}else if state == 17 && char == 'a' {
			state = 21
		} else if state == 18 && char == 'i' {
			state = 22
		} else if state == 19 && char == 'g' {
			state = 20
		} else if state == 20 && char == ' ' {
			state = 23
		} else if state == 21 && char == 't' {
			state = 20
		} else if state == 22 && char == 'l' {
			state = 20
		} else {
			return false
		}
	}
	return state == 20 || state == 23
}

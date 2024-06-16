package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Masukkan kalimat: ")
	reader := bufio.NewReader(os.Stdin)
	kalimat, _ := reader.ReadString('\n')
	kalimat = strings.TrimSpace(kalimat)

	if bacaPredikat(kalimat) {
		fmt.Println("Kalimat termasuk ke predikat")
	} else {
		fmt.Println("Kalimat bukan termasuk ke predikat")
	}
}
//nyoba aja buat predikat, baru ke input memotong

func bacaKalimat(kalimat *string) {
	fmt.Println("Kalimat yang dibaca: ", *kalimat)
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
		switch state {
		case 0:
			if char == 'm' {
				state = 1
			} else {
				return false
			}
		case 1:
			if char == 'e' {
				state = 2
			} else if char == 'a' || char == 'i' || char == 'u' || char == 'o' {
				state = 4
			} else {
				return false
			}
		case 2:
			if char == 'm' {
				state = 3
			} else if char == 'n' || char == 'l' || char == 'g' || char == 'b' {
				state = 4
			} else {
				return false
			}
		case 3:
			if char == 'o' {
				state = 5
			} else if char == 'b' {
				state = 6
			} else if char == 'n' || char == 'l' {
				state = 4
			} else {
				return false
			}
		case 4:
			if char == 'o' || char == 'i' || char == 'a' || char == 'u' {
				state = 5
			} else {
				return false
			}
		case 5:
			if char == 't' || char == 'k' {
				state = 6
			} else if char == 'b' {
				state = 7
			} else if char == 'n' || char == 'l' {
				state = 8
			} else if char == 'a' || char == 'e' || char == 'i' || char == 'u' {
				state = 9
			} else {
				return false
			}
		case 6:
			if char == 'o' || char == 'a' || char == 'e' || char == 'i' || char == 'u' {
				state = 7
			} else {
				return false
			}
		case 7:
			if char == 'n' || char == 'k' {
				state = 8
			} else {
				return false
			}
		case 8:
			if char == 'g' {
				return true
			} else if char == 'i' || char == 'a' || char == 'u' || char == 'e' {
				state = 9
			} else {
				return false
			}
		case 9:
			if char == 'h' || char == 'l' {
				return true
			} else if char == 'n' {
				state = 10
			} else if char == 't' {
				state = 11
			} else {
				return false
			}
		case 10:
			if char == 'g' {
				return true
			} else {
				return false
			}
		case 11:
			if char == 'e' {
				state = 12
			} else if char == 'k' {
				state = 13
			} else {
				return false
			}
		case 12:
			if char == 'h' {
				return true
			} else {
				return false
			}
		case 13:
			if char == 'i' {
				state = 14
			} else {
				return false
			}
		case 14:
			if char == 'n' {
				state = 15
			} else {
				return false
			}
		case 15:
			if char == 'g' {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	}
	return false
}

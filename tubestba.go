package main

import (
	"fmt"
)

type infotype byte
type stack struct {
	info [100]infotype
	top  int
}

func main() {
	var kalimat string
	fmt.Print("Masukkan kalimat: ")
	bacaKalimat(&kalimat)
	if tokenRecognizer(kalimat) {
		fmt.Println("KATA DALAM KALIMAT VALID")
		if parser(kalimat) {
			fmt.Println("STRUKTUR KALIMAT VALID")
		} else {
			fmt.Println("STRUKTUR KALIMAT TIDAK VALID")
		}
	} else {
		fmt.Println("KATA DALAM KALIMAT TIDAK VALID")
	}
}

func bacaKalimat(kalimat *string) {
	var input byte
	fmt.Scanf("%c", &input)
	for input != '\r' {
		*kalimat = *kalimat + string(rune(input))
		fmt.Scanf("%c", &input)
	}
}

func tokenRecognizer(kalimat string) bool {
	var struktur string
	var check string
	var i int
	var j int = -1
	var before string
	for i < len(kalimat) {
		for j = i; j < len(kalimat) && kalimat[j] != ' '; j++ {
			check = check + string(rune(kalimat[j]))
		}
		before = struktur
		bacaSubjek(check, &struktur)
		bacaPredikat(check, &struktur)
		bacaObjek(check, &struktur)
		bacaKeterangan(check, &struktur)
		i = j + 1
		check = ""
		if before == struktur {
			return false
		}
	}
	return true
}

func bacaSubjek(kalimat string, struktur *string) {
	var statement byte
	var point int
	var error bool = false
	for i := 0; i < len(kalimat); i++ {
		statement = kalimat[i]
		if point == 0 {
			if statement == 'a' {
				point = 1
			} else if statement == 'k' {
				point = 2
			} else if statement == 'd' {
				point = 3
			} else if statement == 'm' {
				point = 4
			} else {
				break
			}
		} else if point == 1 && statement == 'k' {
			point = 5
		} else if point == 2 && statement == 'a' {
			point = 6
		} else if point == 3 && statement == 'i' {
			point = 7
		} else if point == 4 && statement == 'e' {
			point = 8
		} else if point == 5 && statement == 'u' {
			point = 13
		} else if point == 6 && statement == 'm' {
			point = 9
		} else if point == 7 && statement == 'a' {
			point = 13
		} else if point == 8 && statement == 'r' {
			point = 10
		} else if point == 9 && (statement == 'u' || statement == 'i') {
			point = 13
		} else if point == 10 && statement == 'e' {
			point = 11
		} else if point == 11 && statement == 'k' {
			point = 12
		} else if point == 12 && statement == 'a' {
			point = 13
		} else if point == 13 {
			if statement == ' ' {
				point = 14
			} else {
				error = true
				break
			}
		} else if point == 14 {
			if statement == ' ' {
				point = 14
			} else if statement == 'a' {
				point = 1
			} else if statement == 'k' {
				point = 2
			} else if statement == 'd' {
				point = 3
			} else if statement == 'm' {
				point = 4
			} else {
				error = true
				break
			}
		} else {
			break
		}
	}
	if (point == 13 || point == 14) && error == false {
		*struktur = *struktur + "S"
	}
}

func bacaPredikat(kalimat string, struktur *string) {
	var statement byte
	var point int
	var error bool = false
	for i := 0; i < len(kalimat); i++ {
		statement = kalimat[i]
		if point == 0 && statement == 'm' {
			point = 1
		} else if point == 1 && statement == 'e' {
			point = 2
		} else if point == 2 {
			if statement == 'm' {
				point = 3
			} else if statement == 'l' {
				point = 4
			} else if statement == 'n' {
				point = 5
			} else {
				break
			}
		} else if point == 3 {
			if statement == 'o' {
				point = 6
			} else if statement == 'b' {
				point = 7
			} else {
				break
			}
		} else if point == 4 && statement == 'i' {
			point = 8
		} else if point == 5 && statement == 'g' {
			point = 9
		} else if point == 6 && statement == 't' {
			point = 10
		} else if point == 7 && statement == 'u' {
			point = 11
		} else if point == 8 && statement == 'h' {
			point = 11
		} else if point == 9 {
			if statement == 'a' {
				point = 13
			} else if statement == 'i' {
				point = 12
			} else {
				break
			}
		} else if point == 10 && statement == 'o' {
			point = 14
		} else if point == 11 && statement == 'a' {
			point = 15
		} else if point == 12 && statement == 'k' {
			point = 16
		} else if point == 13 && statement == 'm' {
			point = 17
		} else if point == 14 && statement == 'n' {
			point = 18
		} else if point == 15 && statement == 't' {
			point = 22
		} else if point == 16 && statement == 'a' {
			point = 19
		} else if point == 17 && statement == 'b' {
			point = 20
		} else if point == 18 && statement == 'g' {
			point = 22
		} else if point == 19 && statement == 't' {
			point = 22
		} else if point == 20 && statement == 'i' {
			point = 21
		} else if point == 21 && statement == 'l' {
			point = 22
		} else if point == 22 {
			if statement == ' ' {
				point = 23
			} else {
				error = true
				break
			}
		} else if point == 23 {
			if statement == ' ' {
				point = 23
			} else if statement == 'm' {
				point = 1
			} else {
				error = true
				break
			}
		} else {
			break
		}
	}
	if (point == 22 || point == 23) && error == false {
		*struktur = *struktur + "P"
	}
}

func bacaObjek(kalimat string, struktur *string) {
	var statement byte
	var point int
	var error bool = false
	for i := 0; i < len(kalimat); i++ {
		statement = kalimat[i]
		if point == 0 {
			if statement == 'k' {
				point = 1
			} else if statement == 'b' {
				point = 2
			} else if statement == 'p' {
				point = 3
			} else {
				break
			}
		} else if point == 1 && statement == 'a' {
			point = 4
		} else if point == 2 && statement == 'e' {
			point = 5
		} else if point == 3 && statement == 'a' {
			point = 6
		} else if point == 4 && statement == 'y' {
			point = 7
		} else if point == 5 {
			if statement == 's' {
				point = 8
			} else if statement == 'n' {
				point = 9
			} else {
				break
			}
		} else if point == 6 {
			if statement == 'p' {
				point = 10
			} else if statement == 't' {
				point = 11
			} else {
				break
			}
		} else if point == 7 && statement == 'u' {
			point = 18
		} else if point == 8 && statement == 'i' {
			point = 18
		} else if point == 9 && statement == 'd' {
			point = 12
		} else if point == 10 && statement == 'a' {
			point = 13
		} else if point == 11 && statement == 'u' {
			point = 14
		} else if point == 12 && statement == 'e' {
			point = 15
		} else if point == 13 && statement == 'n' {
			point = 18
		} else if point == 14 && statement == 'n' {
			point = 16
		} else if point == 15 && statement == 'r' {
			point = 17
		} else if point == 16 && statement == 'g' {
			point = 18
		} else if point == 17 && statement == 'a' {
			point = 18
		} else if point == 18 {
			if statement == ' ' {
				point = 19
			} else {
				error = true
				break
			}
		} else if point == 19 {
			if statement == ' ' {
				point = 19
			} else if statement == 'k' {
				point = 1
			} else if statement == 'b' {
				point = 2
			} else if statement == 'p' {
				point = 3
			} else {
				error = true
				break
			}
		} else {
			break
		}
	}
	if (point == 18 || point == 19) && error == false {
		*struktur = *struktur + "O"
	}
}

func bacaKeterangan(kalimat string, struktur *string) {
	var statement byte
	var point int
	var error bool = false
	for i := 0; i < len(kalimat); i++ {
		statement = kalimat[i]
		if point == 0 {
			if statement == 's' {
				point = 1
			} else if statement == 'b' {
				point = 2
			} else if statement == 't' {
				point = 3
			} else if statement == 'n' {
				point = 4
			} else {
				break
			}
		} else if point == 1 && statement == 'e' {
			point = 5
		} else if point == 2 && statement == 'e' {
			point = 6
		} else if point == 3 && statement == 'a' {
			point = 7
		} else if point == 4 && statement == 'a' {
			point = 8
		} else if point == 5 {
			if statement == 'k' {
				point = 9
			} else if statement == 'm' {
				point = 10
			} else {
				break
			}
		} else if point == 6 && statement == 's' {
			point = 11
		} else if point == 7 && statement == 'd' {
			point = 12
		} else if point == 8 && statement == 'n' {
			point = 13
		} else if point == 9 && statement == 'a' {
			point = 14
		} else if point == 10 && statement == 'a' {
			point = 15
		} else if point == 11 && statement == 'o' {
			point = 16
		} else if point == 12 && statement == 'i' {
			point = 22
		} else if point == 13 && statement == 't' {
			point = 12
		} else if point == 14 && statement == 'r' {
			point = 17
		} else if point == 15 && statement == 'l' {
			point = 18
		} else if point == 16 && statement == 'k' {
			point = 22
		} else if point == 17 && statement == 'a' {
			point = 19
		} else if point == 18 && statement == 'a' {
			point = 20
		} else if point == 19 && statement == 'n' {
			point = 21
		} else if point == 20 && statement == 'm' {
			point = 22
		} else if point == 21 && statement == 'g' {
			point = 22
		} else if point == 22 {
			if statement == ' ' {
				point = 23
			} else {
				error = true
				break
			}
		} else if point == 23 {
			if statement == ' ' {
				point = 23
			} else if statement == 'e' {
				point = 1
			} else if statement == 'b' {
				point = 2
			} else if statement == 't' {
				point = 3
			} else if statement == 'n' {
				point = 4
			} else {
				error = true
				break
			}
		} else {
			break
		}
	}
	if (point == 22 || point == 23) && error == false {
		*struktur = *struktur + "K"
	}
}

func parser(kata string) bool {
	var S stack
	createStack(&S)
	var atas infotype
	var i int = -1
	var error bool = false
	push(&S, '#')
	push(&S, 'I')
	for S.info[S.top] != '#' {
		atas = S.info[S.top]
		if atas == 'I' {
			pop(&S)
			multipush(&S, "SPOK")
		}
		if atas == 'S' {
			pop(&S)
			if kata[i] == 'a' {
				multipush(&S, "aku")
			} else if kata[i] == 'k' && kata[i+1] == 'a' && kata[i+2] == 'm' {
				if kata[i+3] == 'i' {
					multipush(&S, "kami")
				} else if kata[i+3] == 'u' {
					multipush(&S, "kamu")
				}
			} else if kata[i] == 'd' {
				multipush(&S, "dia")
			} else if kata[i] == 'm' && kata[i+1] == 'e' && kata[i+2] == 'r' {
				multipush(&S, "mereka")
			} else {
				multipush(&S, "S")
				break
			}
		}
		if atas == 'P' {
			pop(&S)
			if kata[i] == 'm' && kata[i+1] == 'e' {
				if kata[i+2] == 'm' {
					if kata[i+3] == 'o' {
						multipush(&S, "memotong")
					} else if kata[i+3] == 'b' {
						multipush(&S, "membuat")
					}
				} else if kata[i+2] == 'n' && kata[i+3] == 'g' {
					if kata[i+4] == 'i' {
						multipush(&S, "mengikat")
					} else if kata[i+4] == 'a' {
						multipush(&S, "mengambil")
					}
				} else if kata[i+2] == 'l' {
					multipush(&S, "melihat")
				}
			} else {
				multipush(&S, "P")
				break
			}
		}
		if atas == 'O' {
			pop(&S)
			if kata[i] == 'k' && kata[i+1] == 'a' && kata[i+2] == 'y' {
				multipush(&S, "kayu")
			} else if kata[i] == 'b' && kata[i+1] == 'e' && kata[i+2] == 'n' {
				multipush(&S, "bendera")
			} else if kata[i] == 'b' && kata[i+1] == 'e' && kata[i+2] == 's' && kata[i+3] == 'i' {
				multipush(&S, "besi")
			} else if kata[i] == 'p' && kata[i+1] == 'a' {
				if kata[i+2] == 'p' {
					multipush(&S, "papan")
				} else if kata[i+2] == 't' {
					multipush(&S, "patung")
				}
			} else if i != len(kata)-1 {
				i = i - 1
			}
		}
		if atas == 'K' {
			pop(&S)
			if kata[i] == 's' && kata[i+1] == 'e' {
				if kata[i+2] == 'k' {
					multipush(&S, "sekarang")
				} else if kata[i+2] == 'm' {
					multipush(&S, "semalam")
				}
			} else if kata[i] == 'b' && kata[i+1] == 'e' && kata[i+2] == 's' {
				multipush(&S, "besok")
			} else if kata[i] == 't' {
				multipush(&S, "tadi")
			} else if kata[i] == 'n' && kata[i+1] == 'a' {
				multipush(&S, "nanti")
			}
		}
		if atas == 'a' {
			pop(&S)
		}
		if atas == 'b' {
			pop(&S)
		}
		if atas == 'd' {
			pop(&S)
		}
		if atas == 'e' {
			pop(&S)
		}
		if atas == 'g' {
			pop(&S)
		}
		if atas == 'h' {
			pop(&S)
		}
		if atas == 'i' {
			pop(&S)
		}
		if atas == 'k' {
			pop(&S)
		}
		if atas == 'l' {
			pop(&S)
		}
		if atas == 'm' {
			pop(&S)
		}
		if atas == 'n' {
			pop(&S)
		}
		if atas == 'o' {
			pop(&S)
		}
		if atas == 'p' {
			pop(&S)
		}
		if atas == 'r' {
			pop(&S)
		}
		if atas == 's' {
			pop(&S)
		}
		if atas == 't' {
			pop(&S)
		}
		if atas == 'u' {
			pop(&S)
		}
		if atas == 'w' {
			pop(&S)
		}
		if atas == 'y' {
			pop(&S)
		}
		if i < len(kata)-1 {
			i = i + 1
		}
		if S.info[S.top] == '#' && i != len(kata)-1 {
			error = true
		}
	}
	if pop(&S) == '#' && !error {
		return true
	}
	return false
}

func createStack(S *stack) {
	S.top = -1
}

func push(S *stack, x infotype) {
	S.top = S.top + 1
	S.info[S.top] = x
}

func multipush(S *stack, x string) {
	for i := len(x) - 1; i >= 0; i-- {
		S.top = S.top + 1
		S.info[S.top] = infotype(x[i])
	}
}

func pop(S *stack) infotype {
	var x infotype
	x = S.info[S.top]
	S.top = S.top - 1
	return x
}

func printInfo(S stack) {
	fmt.Print("TOP -> ")
	for i := 0; i <= S.top; i++ {
		fmt.Print(string(rune(S.info[i])), " ")
	}
	fmt.Println()
}

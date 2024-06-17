package main

import "fmt"

func main() {
	var kalimat string
	var struktur string
	var check string
	var i int
	var j int = -1
	bacaKalimat(&kalimat)
	for i < len(kalimat) {
		for j = i; j < len(kalimat) && kalimat[j] != ' '; j++ {
			fmt.Println(i, j)
			check = check + string(rune(kalimat[j]))
		}
		fmt.Println(check)
		bacaSubjek(check, &struktur)
		bacaPredikat(check, &struktur)
		bacaObjek(check, &struktur)
		bacaKeterangan(check, &struktur)
		i = j + 1
		check = ""
	}
	fmt.Println(struktur)
}

func bacaKalimat(kalimat *string) {
	var input byte
	fmt.Scanf("%c", &input)
	for input != '\r' {
		*kalimat = *kalimat + string(rune(input))
		fmt.Scanf("%c", &input)
	}
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
			if statement == 'p' {
				point = 1
			} else if statement == 'd' {
				point = 2
			}else{
				break
			}
		}else if point == 1 && statement == 'a' {
			point = 3
		}else if point == 2 && statement == 'i' {
			point = 4
		}else if point == 3 && statement == 'd' {
			point = 5
		}else if point == 4 && statement == ' ' {
			point = 6
		}else if point == 5 && statement == 'a' {
			point = 7
		}else if point == 6 {
			if statement == 'r' {
				point = 8
			}else if statement == 'l' {
				point = 9
			}else if statement == 'h' {
				point = 10
			}else{
				break
			}
		}else if point == 7 && statement == ' ' {
			point = 11
		}else if point == 8 && statement == 'u' {
			point = 12
		}else if point == 9 && statement == 'a' {
			point = 13
		}else if point == 10 && statement == 'a' {
			point = 14
		}else if point == 11 && statement == 'p' {
			point = 24
		}else if point == 12 && statement == 'm' {
			point = 15
		}else if point == 13 && statement == 'p' {
			point = 16
		}else if point == 14 && statement == 'l' {
			point = 17
		}else if point == 15 && statement == 'a' {
			point = 18
		}else if point == 16 && statement == 'a' {
			point = 19
		}else if point == 17 && statement == 'a' {
			point = 20
		}else if point == 18 && statement == 'h'{
			point = 36
		}else if point == 19 && statement == 'n' {
			point = 21
		}else if point == 20 && statement == 'm' {
			point = 22
		}else if point == 21 && statement == 'g' {
			point = 22
		}else if point == 22 && statement == 'a' {
			point = 23
		}else if point == 23 && statement == 'n' {
			point = 36
		}else if point == 24 && statement == 'a' {
			point = 26
		}else if point == 25 && statement == 'a' {
			point = 27
		}else if point == 26 && statement =
	}
	if (point == 36 || point == 37) && error == false {
		*struktur = *struktur + "O"
	}
}

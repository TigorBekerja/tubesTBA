package main

import "fmt"

type arrByte [5]byte

func main() {
	var kalimat string

	bacaKalimat(&kalimat)
}

func bacaKalimat(kalimat *string) {
	var statement byte
	var point int
	var error bool = false

	fmt.Scanf("%c", &statement)
	fmt.Println(point)
	for statement != '\r' {
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
		*kalimat = *kalimat + string(rune(statement)) //masukin satu satu ke string
		fmt.Scanf("%c", &statement)
		fmt.Println(point)
	}
	if (point == 13 || point == 14) && error == false {
		fmt.Printf("kalimat termasuk ke subjek")
	} else {
		fmt.Printf("kalimat adalah bukan termasuk ke subjek")
	}
}

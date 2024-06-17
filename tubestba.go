package main

import (
	"fmt"
)

type arrByte [5]byte

func main() {
	var kalimat string
	fmt.Print("Masukkan Kalimat : ")
	bacaSubjek(&kalimat)
	bacaPredikat(&kalimat)
	bacaKeterangan(&kalimat)
	fmt.Println("\nHasil : ", kalimat)
}

//SUBJEK
func bacaSubjek(kalimat *string) {
	var statement byte
	var point int
	var error bool = false
	fmt.Scanf("%c", &statement)
	fmt.Print("\nSUBJEK : ")
	for statement != ' ' {
		fmt.Print(point, " ")
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
		//*kalimat = *kalimat + string(rune(statement)) //masukin satu satu ke string
		fmt.Scanf("%c", &statement)
	}
	if (point == 13 || point == 14) && error == false {
		*kalimat = *kalimat + "S - "
	}
}

// //PREDIKAT
func bacaPredikat(kalimat *string) {
	var statement byte
	var point int
	var error bool = false
	fmt.Print("\nPREDIKAT : ")
	fmt.Scanf("%c", &statement)
	for statement != ' ' {
		fmt.Print(point, " ")
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
		//*kalimat = *kalimat + string(rune(statement)) //masukin satu satu ke string
		fmt.Scanf("%c", &statement)
	}
	if (point == 22 || point == 23) && error == false {
		*kalimat = *kalimat + "P - "
	}
}

func bacaKeterangan(kalimat *string) {
	var statement byte
	var point int
	var error bool = false
	fmt.Print(" \nKETERANGAN : ")
	fmt.Scanf("%c", &statement)
	for statement > 1 {
		fmt.Print(point, " ")
		if point == 0 {
			if statement == 'p'{
				point = 1
			}else if statement == 'd'{
				point = 2
			}else {
				break
			}
		} else if point == 1 && statement == 'd'{
			point = 2
		} else if point == 1 && statement == 'a'{
			point = 3
		} else if point == 2 && statement == 'i'{
			point = 4
		} else if point == 3 && statement == 'd'{
			point = 5
		} else if point == 5 && statement == 'a'{
			point = 24
		} else if point == 4 && statement ==' '{
			point = 6
		} else if point == 6 {
			if statement == ' '{
				point = 6
			} else if statement == 'r'{
				point = 7
			} else if statement == 'l'{
				point = 8
			} else if statement == 'h'{
				point = 9
			} else {
				error = true
				break
			}
		} else if point == 7 && statement == 'u'{
			point = 10
		} else if point == 8 && statement == 'a'{
			point = 11
		} else if point == 9 && statement == 'a'{
		 	point = 12
		} else if point == 10 && statement == 'm'{
			point = 13
		} else if point == 11 && statement == 'p'{
		 	point = 14
		} else if point == 12 && statement == 'l'{
			point = 15
		} else if point == 13 && statement == 'a'{
			point = 16
		} else if point == 14 && statement == 'a'{
			point = 17
		} else if point == 15 && statement == 'a'{
			point = 18
		} else if point == 16 && statement == 'h'{
			point = 22
		} else if point == 17 && statement == 'n'{
			point = 19
		} else if point == 18 && statement == 'm'{
			point = 20
		} else if point == 19 && statement == 'g'{
			point = 20
		} else if point == 20 && statement == 'a'{
			point = 21
		} else if point == 21 && statement == 'n'{
			point = 22
		} else if point == 22{
			if statement == ' '{
				point = 23
			} else {
				// error = true
				break
			}
		} else if point == 23 {
			if statement == ' '{
				point = 23
			}else if statement == 'p' {
				point = 1
			} else if statement == 'd' {
				point = 2
			} else {
				error = true
				break
			}
		} else if point == 24 && statement == ' '{
			point = 25
		} else if point == 25{
			if statement == 'p'{
				point = 26
			} else if statement == 'm'{
				point = 29
			} else {
				error = true
				break
			}

		} else if point == 26 && statement == 'a'{
			point = 27
		} else if point == 27 && statement == 'g'{
			point = 28
		} else if point == 28 && statement == 'i'{
			point = 33
		} else if point == 29 && statement == 'a'{
			point = 30
		} else if point == 30 && statement == 'l'{
			point = 31
		} else if point == 31 && statement == 'a'{
			point = 32
		} else if point == 32 && statement == 'm'{
			point = 33
		} else if point == 33 && statement == ' '{
			point = 34
		} else if point == 34 && statement == 'h'{
			point = 35
		} else if point == 35 && statement == 'a'{
			point = 36
		} else if point == 36 && statement == 'r'{
			point = 37
		} else if point == 37 && statement == 'i'{
			point = 22
		} else {
			break
		}
		fmt.Scanf("%c", &statement)
	
	}
	
	if (point == 22 || point == 23) && error == false {
		*kalimat = *kalimat + "K"
	}
}
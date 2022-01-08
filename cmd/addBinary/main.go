package main

import "fmt"

func main() {
	a := "11"
	b := "1"

	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {

}

/*
str[] a
str[] b
str[] long
str[] short

if a.length < b.length {
	short = a
	long = b
} else {
	short = b
	long = a
}

int difference = long.length - short.length

append(short, )











str[] a = "10"
str[] b = "01"
str[] long
str[] short
str[] res

if a.length < b.length {
	short = a
	long = b
} else {
	short = b
	long = a
}

bool carry = false
for i = long.length - 1; i > 0; i-- {

	if long[i] == '1' && short[i] == '1' {
		res[i] = 0
		carry = true
	} else if (long[i] == '1' && short[i] == '0') || (long[i] == '0' && short[i] == '1) {
		if(carry) {
			res[i] = 0
		}
	}
}











*/

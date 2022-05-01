package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var g, p int
	_, err := fmt.Scanf("g is %d and p is %d\n", &g, &p)
	if err != nil {
		return
	}
	fmt.Println("OK")

	randomB := rand.Intn(p - 1)
	b := compute(g, randomB, p)

	var a int
	_, err = fmt.Scanf("A is %d\n", &a)
	if err != nil {
		return
	}

	s := compute(a, randomB, p) % 26
	fmt.Printf("B is %d\n", b)

	fmt.Println(encrypt("Will you marry me?", s))

	var response string
	_, err = fmt.Scanf("%s", &response)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	switch decrypt(response, s) {
	case "Yeah,":
		fmt.Println(encrypt("Great!", s))
	case "Let's":
		fmt.Println(encrypt("What a pity!", s))
	}
}

func encrypt(text string, shiftNumber int) string {
	return cipher(text, 1, shiftNumber)
}

func decrypt(text string, shiftNumber int) string {
	return cipher(text, -1, shiftNumber)
}

func cipher(text string, direction int, shiftNumber int) string {

	shift, offset := rune(shiftNumber), rune(26)

	runes := []rune(text)

	for index, char := range runes {
		switch direction {
		case -1:
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1:
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		runes[index] = char
	}

	return string(runes)
}

func compute(a int, b int, p int) int {
	r := 1
	for b > 0 {
		r = r * a % p
		b--
	}
	return r
}

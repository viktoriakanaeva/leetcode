package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 9. Palindrome Number
	//Given an integer x, return true if x is a palindrome, and false otherwise.

	fmt.Println("put a number")
	var n int
	fmt.Scan(&n)

	fmt.Println(palindrome(n))

}

func palindrome(n int) bool {

	nstr := strconv.Itoa(n)
	var nstr1, nstr2 string

	if len(nstr)%2 == 0 {
		nstr1 = nstr[:len(nstr)/2]

		nstr2 = nstr[len(nstr)/2:]
	} else {
		nstr1 = nstr[:len(nstr)/2]

		nstr2 = nstr[len(nstr)/2+1:]
	}
	fmt.Println(nstr1, nstr2)

	b := []byte(nstr1)

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	fmt.Println(string(b))

	if strings.Compare(string(b), nstr2) == 0 {
		return true
	}
	return false
}

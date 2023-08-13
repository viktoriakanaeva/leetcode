/*125. Valid Palindrome
A  phrase is a palindrome if, after converting all uppercase letters i lowercase letters and removing all non-alphanumeric
characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "I, man, am regal—a German am I"

	fmt.Println(changeOfStr(s))
}

func changeOfStr(s string) bool {

	s = strings.ToLower(s)

	byteStr := []byte(s)
	var c []byte
	var nstr1, nstr2 string

	for _, v := range byteStr {

		if len(s) == 1 {
			return true
		} else if len(s) != 1 && (v > 96 && v < 123 || v > 47 && v < 58) {
			c = append(c, v)
		}
	}
	if len(c) == 1 {
		return true
	}
	s = string(c)

	nstr1 = s[:len(s)/2]

	if len(s)%2 == 1 {
		nstr2 = s[len(s)/2+1:]
	} else {
		nstr2 = s[len(s)/2:]
	}

	b1 := []byte(nstr1)
	for i, j := 0, len(b1)-1; i < j; i, j = i+1, j-1 {
		b1[i], b1[j] = b1[j], b1[i]

	}

	if strings.Compare(string(b1), nstr2) == 0 {
		return true
	} else {
		return false
	}
}

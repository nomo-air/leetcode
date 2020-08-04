package _345

import (
	"fmt"
	"testing"
)

func reverseVowels(s string) string {
	vowelMap := map[byte]byte{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	sBytes := []byte(s)

	l := 0
	r := len(sBytes) - 1

	for r > l {
		if vowelMap[sBytes[r]] == 1 && vowelMap[sBytes[l]] == 1 {
			sBytes[r], sBytes[l] = sBytes[l], sBytes[r]
			l++
			r--
		} else if vowelMap[sBytes[l]] == 1 {
			r--
		} else if vowelMap[sBytes[r]] == 1 {
			l++
		} else {
			l++
			r--
		}

	}

	return string(sBytes)
}

func Test345(t *testing.T) {
	s := reverseVowels("abcdefg")
	fmt.Printf("%+v", s)
}

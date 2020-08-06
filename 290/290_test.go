package _290

import (
	"fmt"
	"strings"
	"testing"
)

func wordPattern(pattern string, str string) bool {
	strList := strings.Split(str, " ")
	if len(strList) != len(pattern) {
		return false
	}
	patternStrMap := make(map[string]string) // a:dog
	strPatternMap := make(map[string]string) // dog:a

	for i := 1; i <= len(strList); i++ {

		strValue := strings.Join(strList[i-1:i], "")
		patternValue := pattern[i-1 : i]

		if currentStr, ok := patternStrMap[patternValue]; ok {
			if currentStr != strValue {
				return false
			}
		} else {

			if currentPatternValue, ok := strPatternMap[strValue]; ok {
				if currentPatternValue != patternValue {
					return false
				}
			}

			patternStrMap[patternValue] = strValue
			strPatternMap[strValue] = patternValue
		}
	}

	return true
}

func Test290(t *testing.T) {
	b := wordPattern("abba", "dog dog dog dog")
	fmt.Printf("%+v", b)
}

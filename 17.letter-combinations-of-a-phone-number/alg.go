package _7_letter_combinations_of_a_phone_number

import (
	"bytes"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	numMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var combinations []string
	backtrack(&combinations, numMap, digits, 0, &bytes.Buffer{})
	return combinations
}

func backtrack(combinations *[]string, numMap map[string]string, digits string, index int, combination *bytes.Buffer) {
	if index == len(digits) {
		*combinations = append(*combinations, combination.String())
	} else {
		digit := string(digits[index])
		letters := numMap[digit]
		for i := 0; i < len(letters); i++ {
			tmp := combination.String()
			combination.WriteByte(letters[i])
			backtrack(combinations, numMap, digits, index+1, combination)
			combination.Reset()
			combination.WriteString(tmp)
		}
	}
}

/*
* @Author: Ivy
* @Date: 2022/4/26 5:08 PM
**/
package main

import (
	"fmt"
)

func main() {
	//s1 := "III"
	//answer1 := romanToInt(s1)
	//fmt.Println(answer1)
	//
	//s2 := "IV"
	//answer2 := romanToInt(s2)
	//fmt.Println(answer2)
	//
	//s3 := "LVIII"
	//answer3 := romanToInt(s3)
	//fmt.Println(answer3)

	s4 := "CMXCIX"
	answer4 := romanToInt(s4)
	fmt.Println(answer4)

}

func romanToInt(s string) int {
	romanAlphabet := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	answer := 0
	for index, i := range s {
		if index < len(s)-1 && romanAlphabet[string(s[index])] < romanAlphabet[string(s[index+1])] {
			answer = answer - 2*romanAlphabet[string(s[index])]
		}
		answer = answer + romanAlphabet[string(i)]
	}

	return answer
}

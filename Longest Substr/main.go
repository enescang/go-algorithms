//Enes Can Güneş - kodlib.com - 2021
//Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import (
	"fmt"
)

func main() {
	var l int
	l = find("CODINGISAWESOME")
	fmt.Println("CODINGISAWESOME => Longest Substring Length: ", l)

	l = find("abcabcbb")
	fmt.Println("abcabcbb => Longest Substring Length: ", l)
}

//Use Window Sliding Algorithm
//Create a a_point that starts with zero(0)
//And this point must go forward NOT backward
//If current char if already seen check the length of substring
//save and compare with totalElement
func find(str string) int {
	strLength := len(str)
	seen := []string{}
	a_point := 0

	totalElement := 0
	for i := 0; i < strLength; i++ {
		currentChar := string(str[i])
		currentCharIndex := indexOf(seen, currentChar)
		if currentCharIndex > -1 {
			if i-a_point > totalElement {
				totalElement = i - a_point
			}
			if a_point <= currentCharIndex+1 {
				a_point = currentCharIndex + 1
				//Change the character because we already check this char
				seen[currentCharIndex] += "*"
			}
		}
		seen = append(seen, currentChar)

		//fmt.Println("Char: ", currentChar, " i: ", i, " a_point: ", a_point)
	}

	//fmt.Println("TOTAL: ", totalElement)
	if totalElement == 0 {
		return len(str)
	}
	return totalElement
}

//Index of any element
//Source: https://stackoverflow.com/questions/8307478/how-to-find-out-element-position-in-slice
func indexOf(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

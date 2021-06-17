package main

import "fmt"

func main() {
	fmt.Println("23 => ", combination("23"))
	fmt.Println("27 => ", combination("27"))
	fmt.Println("234 => ", combination("234"))
	fmt.Println(" => ", combination(""))
}

//Get the first digit's chars
//Write each char combinationLength/len(digits) times
//Then write other digits' char
//Example: 23
//First: 	a a a b b b c c c
//Seconf:	d e f d e f d e f
func combination(digits string) []string {
	myLetters := map[string][]string{}
	myLetters["2"] = []string{"a", "b", "c"}
	myLetters["3"] = []string{"d", "e", "f"}
	myLetters["4"] = []string{"g", "h", "i"}
	myLetters["5"] = []string{"j", "k", "l"}
	myLetters["6"] = []string{"m", "n", "o"}
	myLetters["7"] = []string{"p", "q", "r", "s"}
	myLetters["8"] = []string{"t", "u", "v"}
	myLetters["9"] = []string{"w", "x", "y", "z"}

	result := []string{}
	if len(digits) < 1 {
		return result
	}
	combinationLength := func() int {
		result := 1
		for _, v := range digits {
			if myLetters[string(v)] != nil {
				result *= len(myLetters[string(v)])
			}
		}
		return result
	}
	//fmt.Println("Total combination must be: ", combinationLength())
	for _, v := range digits {
		if myLetters[string(v)] != nil && len(result) != combinationLength() {
			for _, v2 := range myLetters[string(v)] {
				for i := 0; i < combinationLength()/len(myLetters[string(v)]); i++ {
					result = append(result, string(v2))
				}
			}
		} else {
			for i := 0; i < len(result); i++ {
				result[i] += myLetters[string(v)][i%len(myLetters[string(v)])]
			}
		}
	}

	return result
}

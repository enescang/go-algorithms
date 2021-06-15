//Enes Can Güneş - kodlib.com - 2021
package main

import "fmt"

func main() {
	fmt.Println("()", isValid("()"))
	fmt.Println("()[", isValid("()["))
	fmt.Println("()[]{}", isValid("()[]{}"))
	fmt.Println("([)]", isValid("([)]"))
	fmt.Println("{[]}", isValid("{[]}"))

}

func isValid(str string) bool {
	arr := []string{}
	for _, value := range str {
		charValue := string(value)
		if charValue == "(" || charValue == "{" || charValue == "[" {
			arr = append(arr, charValue)
			continue
		}

		if len(arr) > 0 && arr[len(arr)-1] == "(" && charValue == ")" {
			arr = arr[:len(arr)-1]
			continue
		}

		if len(arr) > 0 && arr[len(arr)-1] == "{" && charValue == "}" {
			arr = arr[:len(arr)-1]
			continue
		}

		if len(arr) > 0 && arr[len(arr)-1] == "[" && charValue == "]" {
			arr = arr[:len(arr)-1]
			continue
		}

	}

	// fmt.Println("Len:", len(arr))
	return len(arr) == 0
}

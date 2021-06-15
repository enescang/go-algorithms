//Enes Can Güneş - kodlib.com - 2021
//Problem: https://leetcode.com/problems/shuffle-the-array/
package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []interface{}{2, 5, 1, 3, 4, 7}
	result := shuffleArray(arr, 3)
	fmt.Println("[2, 5, 1, 3, 4, 7] =>", result)

	arr = []interface{}{1, 2, 3, 4, 4, 3, 2, 1}
	result = shuffleArray(arr, 4)
	fmt.Println("[1, 2, 3, 4, 4, 3, 2, 1] =>", result)

	arr = []interface{}{1, 1, 2, 2}
	result = shuffleArray(arr, 2)
	fmt.Println("[1, 1, 2, 2] =>", result)
}

//We can see pattern must be like this
// i => (length + i - n) => make it simple => (i+n)
// 0 => (i + n)
// 1 => (i + n)
// 2 => (i + n)
// ....
func shuffleArray(a interface{}, n int) []interface{} {
	arr := a.([]interface{})
	var arrLength int = len(arr)
	if 2*n != arrLength {
		log.Fatal("Array length must be 2*(n)")
	}

	temp := []interface{}{}
	var flag int
	for i := 0; i < arrLength/2; i++ {
		temp = append(temp, arr[i])
		temp = append(temp, arr[i+n])
		flag++
	}
	return temp
}

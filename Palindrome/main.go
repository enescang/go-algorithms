package main

import "math"

func main() {
	println("123 is palindrome? : ", isPalindrome(123))
	println("121 is palindrome? : ", isPalindrome(121))
	println("737 is palindrome? : ", isPalindrome(737))
	println("1111 is palindrome? : ", isPalindrome(1111))
}

func isPalindrome(num int) bool {
	temp := num
	var result int
	index := totalDigits(num)
	for temp > 0 {
		mod := temp % (10)
		//println("MOD: ", mod)
		result += mod * int(math.Pow(10, float64(index)-1))
		//println("RESULT: ", result)
		temp = temp / 10
		//println("NEXT NUMBER: ", temp)
		index -= 1
	}
	return result-num == 0
}

func totalDigits(num int) int {
	digits := 0
	for num > 0 {
		num /= 10
		digits += 1
	}
	//println("Total Digits:", digits)
	return digits
}

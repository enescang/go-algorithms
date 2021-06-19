//Enes Can Güneş - kodlib.com - 2021
//Problem: https://leetcode.com/problems/restore-ip-addresses/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	restoreIpAddresses("101023")
	restoreIpAddresses("25525511135")
	restoreIpAddresses("0000")
	restoreIpAddresses("1111")
	restoreIpAddresses("010010")
	restoreIpAddresses("192168401")
}

func restoreIpAddresses(IP string) {
	var arr = make([]string, 0)
	a, b, c, d := 3, 3, 3, 3
	//Example: 192.168.40.1
	//aPart: 192 => 3 digits(a)
	//bPart: 168 => 3 digits(b)
	//cPart: 40  => 2 digits(c)
	//dPart: 1   => 1 digit(d)
	fmt.Println("The IP:", IP)

	for i := 0; i < 81; i++ {
		if a+b+c+d == len(IP) {
			parseIp(IP, a, b, c, d, &arr)
		}

		//Create the pattern:
		// 3 3 3 3
		// 3 3 3 2
		// ...
		// 1 1 1 1
		if d != 0 {
			d--
			if d == 0 {
				c--
				d = 3
				if c == 0 {
					b--
					c = 3
					if b == 0 {
						a--
						b = 3
					}
				}
			}
		}

	}
	fmt.Println("Total", len(arr), "IP Found")
	fmt.Println(arr)
	fmt.Println("----------------------------")
}

func parseIp(IP string, a, b, c, d int, arr *[]string) {
	result := ""
	temp := IP
	aPart, temp := temp[0:a], temp[a:]
	bPart, temp := temp[0:b], temp[b:]
	cPart, temp := temp[0:c], temp[c:]
	dPart := temp[0:d]

	if (len(aPart) == 1 && aPart[:1] == "0") || strToInt(aPart) <= 256 && aPart[:1] != "0" {
		result += aPart + "."
	} else {
		return
	}
	if (len(bPart) == 1 && bPart[:1] == "0") || strToInt(bPart) <= 256 && bPart[:1] != "0" {
		result += bPart + "."
	} else {
		return
	}
	if (len(cPart) == 1 && cPart[:1] == "0") || strToInt(cPart) <= 256 && cPart[:1] != "0" {
		result += cPart + "."
	} else {
		return
	}
	if (len(dPart) == 1 && dPart[:1] == "0") || strToInt(dPart) <= 256 && dPart[:1] != "0" {
		result += dPart
	} else {
		return
	}
	*arr = append(*arr, result)
	// fmt.Println(aPart, bPart, cPart, dPart)
}

func strToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

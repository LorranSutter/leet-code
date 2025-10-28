package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)
	lenGcd := GCD(len1, len2)

	if len1 < len2 {
		str1, str2 = str2, str1
	}

	div := string(str2[:lenGcd])
	if divides(str1, div) && divides(str2, div) {
		return div
	}

	return ""
}

func divides(str string, div string) bool {
	lenStr, lenDiv := len(str), len(div)

	shift := lenStr / lenDiv
	for i := range div {
		for j := range shift {
			if str[i+lenDiv*j] != div[i] {
				return false
			}
		}
	}

	return true
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC") == "ABC")
	fmt.Println(gcdOfStrings("ABABAB", "ABAB") == "AB")
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB") == "ABAB")
	fmt.Println(gcdOfStrings("LEET", "CODE") == "")
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX") == "TAUXX")
	fmt.Println(gcdOfStrings("AAAAAAAAA", "AAACCC") == "")
}

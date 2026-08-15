package main

import (
	"strconv"
	"strings"

	"leetcode/utils"
)

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func compareVersion1(version1 string, version2 string) int {
	version1arr := strings.Split(version1, ".")
	version2arr := strings.Split(version2, ".")

	len1 := len(version1arr)
	len2 := len(version2arr)

	swapped := false
	if len1 != len2 {
		if len1 > len2 {
			len1, len2 = len2, len1
			version1arr, version2arr = version2arr, version1arr
			swapped = true
		}
	}

	i := 0
	for ; i < len1; i++ {
		if version1arr[i] != version2arr[i] {
			num1, _ := strconv.Atoi(version1arr[i])
			num2, _ := strconv.Atoi(version2arr[i])
			if num1 > num2 {
				if swapped {
					return -1
				}
				return 1
			}
			if num1 < num2 {
				if swapped {
					return 1
				}
				return -1
			}
		}
	}

	for j := i; j < len2; j++ {
		if num2, _ := strconv.Atoi(version2arr[j]); num2 > 0 {
			if swapped {
				return 1
			}
			return -1
		}
	}

	return 0
}

func compareVersion2(version1 string, version2 string) int {
	version1arr := strings.Split(version1, ".")
	version2arr := strings.Split(version2, ".")

	len1 := len(version1arr)
	len2 := len(version2arr)

	if len1 != len2 {
		if len1 > len2 {
			diffSize := len1 - len2
			zeros := make([]string, diffSize)

			for i := range diffSize {
				zeros[i] = "0"
			}

			version2arr = append(version2arr, zeros...)
		} else {
			diffSize := len2 - len1
			zeros := make([]string, diffSize)

			for i := range diffSize {
				zeros[i] = "0"
			}

			version1arr = append(version1arr, zeros...)
		}
	}

	for i := 0; i < max(len1, len2); i++ {
		if version1arr[i] != version2arr[i] {
			num1, _ := strconv.Atoi(version1arr[i])
			num2, _ := strconv.Atoi(version2arr[i])
			if num1 > num2 {
				return 1
			}
			if num1 < num2 {
				return -1
			}
		}
	}

	return 0
}

func main() {
	utils.RunTests([]utils.TestCase[int]{
		{Input: []string{"1", "1.0.0"}, Got: compareVersion1("1", "1.0.0"), Expected: 0},
		{Input: []string{"1", "1.10"}, Got: compareVersion1("1", "1.10"), Expected: -1},
		{Input: []string{"1.10", "1"}, Got: compareVersion1("1.10", "1"), Expected: 1},
		{Input: []string{"1.2", "1.10"}, Got: compareVersion1("1.2", "1.10"), Expected: -1},
		{Input: []string{"2.2", "1.10"}, Got: compareVersion1("2.2", "1.10"), Expected: 1},
		{Input: []string{"0.2", "1.10"}, Got: compareVersion1("0.2", "1.10"), Expected: -1},
		{Input: []string{"1.01", "1.001"}, Got: compareVersion1("1.01", "1.001"), Expected: 0},
		{Input: []string{"1.0", "1.0.0.0"}, Got: compareVersion1("1.0", "1.0.0.0"), Expected: 0},
		{Input: []string{"1.0", "1.0.0.0.1"}, Got: compareVersion1("1.0", "1.0.0.0.1"), Expected: -1},
		{Input: []string{"1.0.0.0.1", "1.0"}, Got: compareVersion1("1.0.0.0.1", "1.0"), Expected: 1},
		{Input: []string{"7.5.2.4", "7.5.3"}, Got: compareVersion1("7.5.2.4", "7.5.3"), Expected: -1},
		{Input: []string{"7.5.3", "7.5.2.4"}, Got: compareVersion1("7.5.3", "7.5.2.4"), Expected: 1},
	})

	utils.RunTests([]utils.TestCase[int]{
		{Input: []string{"1", "1.0.0"}, Got: compareVersion2("1", "1.0.0"), Expected: 0},
		{Input: []string{"1", "1.10"}, Got: compareVersion2("1", "1.10"), Expected: -1},
		{Input: []string{"1.10", "1"}, Got: compareVersion2("1.10", "1"), Expected: 1},
		{Input: []string{"1.2", "1.10"}, Got: compareVersion2("1.2", "1.10"), Expected: -1},
		{Input: []string{"2.2", "1.10"}, Got: compareVersion2("2.2", "1.10"), Expected: 1},
		{Input: []string{"0.2", "1.10"}, Got: compareVersion2("0.2", "1.10"), Expected: -1},
		{Input: []string{"1.01", "1.001"}, Got: compareVersion2("1.01", "1.001"), Expected: 0},
		{Input: []string{"1.0", "1.0.0.0"}, Got: compareVersion2("1.0", "1.0.0.0"), Expected: 0},
		{Input: []string{"1.0", "1.0.0.0.1"}, Got: compareVersion2("1.0", "1.0.0.0.1"), Expected: -1},
		{Input: []string{"1.0.0.0.1", "1.0"}, Got: compareVersion2("1.0.0.0.1", "1.0"), Expected: 1},
		{Input: []string{"7.5.2.4", "7.5.3"}, Got: compareVersion2("7.5.2.4", "7.5.3"), Expected: -1},
		{Input: []string{"7.5.3", "7.5.2.4"}, Got: compareVersion2("7.5.3", "7.5.2.4"), Expected: 1},
	})
}

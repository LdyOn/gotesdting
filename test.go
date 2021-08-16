package main

import "strings"

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count, ok := counts[word]
		if ok {
			counts[word] = count + 1
		} else {
			counts[word] = 1
		}
	}

	return counts
}
func fibonacci() func() int {
	first := 0
	second := 1
	number := 1
	return func() int {
		if number == 1 {
			number++
			return 0
		}
		if number == 2 {
			number++
			return 1
		}
		result := first + second
		first = second
		second = result
		number++
		return result
	}
}

type IPAddr [4]byte

func (v *IPAddr) String() string {
	var ret string
	ret = ""
	for _, value := range v {
		ret = string(value) + "."
	}
	return ret
}

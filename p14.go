package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	commonPrefix := ""
	switch len(strs) {
	case 0:
		return commonPrefix
	case 1:
		return strs[0]
	}

	i := 0
	firstStr := strs[0]
	if len(firstStr) == 0 {
		return commonPrefix
	}
	for {
		c := 1
		for ; c < len(strs); c++ {
			if i >= len(strs[c]) || i >= len(firstStr) || strs[c][i] != firstStr[i] {
				break
			}
		}
		if c == len(strs) {
			i++
		} else {
			break
		}
	}
	return firstStr[:i]
}

func main() {
	strs := []string{"a", "ac"}
	//strs := []string{"abcd"}
	s := longestCommonPrefix(strs)
	fmt.Println(s)
}

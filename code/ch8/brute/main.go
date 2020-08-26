package main

import "fmt"

func main() {
	str1 := "abcdefgaadskgfhidf"
	str2 := "adskg"
	fmt.Println(StringFind(str1, str2))
}

// 优化
func optimization(str1, str2 string) int {
	length1 := len(str1)
	length2 := len(str2)
	if length1 < 0 || length2 < 0 {
		return -1
	}
	if length1 < length2 {
		return -1
	} else {
		pos := -1
		for i := 0; i < length1-length2; i++ {
			isFind := true
			ipos := i
			for j := 0; j < length2; j++ {
				ch1 := str1[i+j]
				ch2 := str2[j]
				if ch1 != ch2 {
					isFind = false
					break
				}
				if ch1 == str2[0] && j >= 1 {
					ipos = i + j
					i = ipos
					fmt.Println("ipos:", ipos)
					break
				}
			}
			if isFind {
				pos = i
				break
			}
		}
		return pos
	}
}

func StringFind(str1 string, str2 string) int {
	length1 := len(str1)
	length2 := len(str2)
	if length1 < 0 || length2 < 0 {
		return -1
	}
	if length1 < length2 {
		return -1
	} else {
		pos := -1
		for i := 0; i < length1-length2; i++ {
			isFind := true
			for j := 0; j < length2; j++ {
				if str1[i+j] != str2[j] {
					isFind = false
					break
				}
			}
			if isFind {
				pos = i
				break
			}
		}
		return pos
	}
}

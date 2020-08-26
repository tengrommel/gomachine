package main

import "fmt"

// 返回数组，next, abcdefg a, ab , abc, abcd
func GetNextValueArray(sub []byte) (next []int) {
	var (
		length       int = len(sub) // 计算长度
		middle       int            // 中间值
		compareLeft  int            // 左边进行对比
		compareRight int            // 右边进行对比
		matchCount   int            // 对应匹配的长度
	)
	next = make([]int, length) // 获取next长度
	next[1] = 0
	next[0] = 0
	for i := 2; i < length; i++ {
		middle = length / 2
		matchCount = 0 // 匹配长度
		// abcabc
		// ababa
		if i%2 == 0 {
			for j := 0; j <= middle; j++ {
				compareLeft = 0          // 第一个字符
				compareRight = i - 1 - j // 最后一个字符
				for compareLeft <= j {
					if sub[compareLeft] != sub[compareRight] {
						break
					}
					compareRight++
					compareLeft++
				}
				if compareLeft == j+1 {
					matchCount++
				}
			}
			next[i] = matchCount
		} else {

		}
	}
	return next
}

func ReviseNextArray(next []int) []int {
	var length int = len(next)
	for i := 2; i < length; i++ {
		if next[i] == next[next[i]] {
			next[i] = next[next[i]] // next数组的叠加
		}
		next[i] = next[next[i]]
	}
	return next
}

func KMP(content []byte, startIndex, endIndex int, sub []byte) (index int) {
	var (
		next      []int = ReviseNextArray(GetNextValueArray(sub)) // 计算next数组
		subIndex  int   = 0
		subLength int   = len(sub)
	)
	for i := startIndex; i <= endIndex; i++ {
		if content[i] == sub[subIndex] {
			matchStart := i
			for j := subIndex; j <= subIndex; j++ {
				if j == subLength { // 找到了
					return matchStart - subIndex
				}
				if i >= endIndex || content[i] != sub[j] {
					subIndex = next[j]
					break
				}
				i++
			}
		}
	}
	return -1
}

// 图论

func main() {
	str1 := []byte("abcdefgaadskgfhidf")
	str2 := []byte("adskg")
	fmt.Println(KMP(str1, 0, len(str1)-1, str2))
}

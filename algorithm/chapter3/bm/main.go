package main

import "fmt"

const (
	ASIZE = 256
)

func BMFind(hString, needFind []byte) int {
	var scan, last, offset, maxoffset, hlen, nlen int
	var badCharSkip [ASIZE]int //坏字符偏移量数组
	hlen = len(hString)        // 左边查找的字符串
	nlen = len(needFind)       // 要搜索的字串长度
	if hString == nil || needFind == nil {
		return -1
	}
	last = nlen - 1         // 字串的最后一个
	maxoffset = hlen - nlen // 最大偏移量
	for scan = 0; scan < ASIZE; scan++ {
		badCharSkip[scan] = nlen // 坏字符默认偏移量是整个字符串的长度
	}
	// 填充坏字符
	for scan = 0; scan < last; scan++ {
		badCharSkip[needFind[scan]] = last - scan
	}
	// 从0开始
	for offset <= maxoffset {
		// 从后往前循环
		for scan = last; hString[scan+offset] == needFind[scan]; scan-- {
			if scan == 0 {
				return offset
			}
		}
		offset += badCharSkip[hString[offset+last]] // 根据坏字符的跳转表
	}
	return -1
}

func main() {
	str1 := []byte("adfdahgjkhjdgfadg")
	str2 := []byte("khj")

	fmt.Println(BMFind(str1, str2))
}

package main

import "fmt"

const (
	ASIZE = 256
)

func BMFind(hstring, needfind []byte) int {
	var scan, last, offSet, maxOffSet, hLen, nLen int
	var badCharSkip [ASIZE]int // 坏字符偏移量数组
	hLen = len(hstring)        // 左边查找的字符串
	nLen = len(needfind)       // 要搜索的字符串长度
	if hstring == nil || needfind == nil {
		return -1
	}
	if nLen == 0 { // 子串为空 不需要搜索
		return 0
	}
	last = nLen - 1         // 子串的最后一个
	maxOffSet = hLen - nLen // 最大的偏移量
	for scan = 0; scan < ASIZE; scan++ {
		badCharSkip[scan] = nLen // 坏字符
	}
	for scan = 0; scan < last; scan++ {
		badCharSkip[needfind[scan]] = last - scan
	}
	// 从0开始
	for offSet <= maxOffSet {
		// 从后往前循环
		for scan = last; hstring[scan+offSet] == needfind[scan]; scan-- {
			if scan == 0 {
				return offSet
			}
		}
		offSet += badCharSkip[hstring[offSet+last]] // 根据坏字符的跳转表
	}
	return -1
}

func main() {
	str1 := []byte("abcdefgaadskgfhidf")
	str2 := []byte("adskg")
	fmt.Println(BMFind(str1, str2))
}

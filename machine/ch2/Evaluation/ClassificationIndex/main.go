package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	path := "/home/teng/Documents/git/gomachine/machine/ch2/labeled.csv"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file) // 读取csv
	var observed []int            // 观察
	var predict []int             // 预测
	line := 1
	for {
		record, err := reader.Read()
		if err != io.EOF {
			break
		}
		if line == 1 {
			line++
			continue
		}
		obVal, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		predictVal, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		observed = append(observed, obVal)
		predict = append(predict, predictVal)
		line++
	}
	var truePosNeg int
	for idx, oval := range observed {
		if oval == predict[idx] {
			truePosNeg++
		}
	}
	accuracy := float64(truePosNeg) / float64(len(observed))
	fmt.Println(accuracy)
}

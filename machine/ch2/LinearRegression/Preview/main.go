package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

// 预览
func main() {
	path := "/home/teng/Documents/git/gomachine/machine/ch2/LinearRegression/Advertising.csv"
	AdFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer AdFile.Close()
	ADDF := dataframe.ReadCSV(AdFile)
	fmt.Println(ADDF.Describe())
	fmt.Println(ADDF)
}

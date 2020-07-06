package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
)


func main() {
	fmt.Println("Enter the xkcd comic number:")
	var num int
	fmt.Scanf("%d", &num)
	url := fmt.Sprintf("https://xkcd.com/%d", num)
	resp, _ := soup.Get(url)
	doc := soup.HTMLParse(resp)
	title := doc.Find("div", "id", "ctitle").Text()
	fmt.Println("Title of the comic :", title)
}

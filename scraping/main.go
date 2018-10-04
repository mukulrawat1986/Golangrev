package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func postScrape(url string, res *[]string) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		title, _ := linkTag.Attr("title")
		title = strings.Split(strings.TrimSpace(title), " ")[0]
		if len(title) != 0 && title != "List" {
			//s := fmt.Sprint("\"", title, "\"")
			*res = append(*res, title)
		}
	})
	//	fmt.Println(res)
}

func main() {
	res := []string{}
	url1 := "https://www.nordicnames.de/wiki/List_of_approved_Icelandic_male_names"
	url2 := "https://www.nordicnames.de/wiki/List_of_approved_Icelandic_female_names"
	postScrape(url1, &res)
	postScrape(url2, &res)

	// sort the slice
	sort.Strings(res)
	//echo(res)

	// create last names
	lres := []string{}
	lastName(res, &lres)
	echo(lres)
}

func lastName(res []string, lres *[]string) {

	for _, name := range res {
		s := fmt.Sprint("\"", name, "sson", "\"")
		s1 := fmt.Sprint("\"", name, "sdóttir", "\"")
		*lres = append(*lres, s, s1)
	}

}

func echo(res []string) {

	s, sep := "", ""

	for _, arg := range res {
		s += sep + arg
		sep = ", "
	}
	fmt.Println(s)
}

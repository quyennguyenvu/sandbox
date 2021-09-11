package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	file, err := os.OpenFile("phone_numbers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	chn := make(chan string)
	count := 0

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	go crawling(".nav-previous > a", "prev", chn)
	go crawling(".nav-next > a", "next", chn)

	for {
		v, ok := <-chn
		if !ok {
			break
		}
		datawriter.WriteString(v + "\n")
		count++
		if count == 100 {
			datawriter.Flush()
			fmt.Println("----------------- FLUSH -----------------")
			count = 0
		}
	}

	datawriter.Flush()
	file.Close()
	fmt.Println("Terminating Crawling")
}

func crawling(hrefClass, name string, chn chan<- string) {
	c := colly.NewCollector()
	count := 0
	c.OnHTML(".article-inner", func(e *colly.HTMLElement) {
		number := strings.TrimPrefix(e.ChildText("h1.entry-title"), "Khach Hang – ")
		chn <- number

		next_href := e.ChildAttr(hrefClass, "href")
		next_href = e.Request.AbsoluteURL(next_href)
		count++
		e.Request.Visit(next_href)
	})
	c.Wait()

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("%v. %v Visiting %v\n", count, name, r.URL)
	})

	c.Visit("https://www.pharmacity.vn/prescription_history/khach-hang-0347708703-3/")

}

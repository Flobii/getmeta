package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/flobii/getmeta/stringsplus"
	"github.com/gocolly/colly"
)

type metaElem struct {
	Name    string
	Content string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: getmeta [url]\nCLI mode: getmeta --cli")
		return
	}
	if os.Args[1] == "--cli" {
		cliURL()
		return
	}
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println("Wrong number of arguments")
		return
	}
	singleURL(os.Args[1])
}

func cliURL() {
	const message string = "Enter URL to get meta. To exit press `Ctrl + D`"
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("URL: ")
		url, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				return
			} else {
				fmt.Println(err)
			}
		}
		url = strings.Replace(url, "\n", "", -1)
		if url != "" {
			fmt.Println(strings.Repeat("─", utf8.RuneCountInString(url)+5))
			singleURL(url)
			fmt.Println()
		} else {
			fmt.Print()
		}
	}
}

func singleURL(url string) {
	c := colly.NewCollector()
	c.OnHTML("html", handleHTML)

	c.Visit(url)
}

func handleHTML(e *colly.HTMLElement) {
	names := []string{}
	contents := []string{}
	e.ForEach("meta", func(i int, e *colly.HTMLElement) {
		name, content := e.Attr("name"), e.Attr("content")
		if name != "" && content != "" {
			names = append(names, name)
			contents = append(contents, content)
		}
	})

	meta := formatMeta(names, contents)
	bold := color.New(color.Bold)
	for _, item := range meta {
		bold.Print(item.Name)
		fmt.Println(item.Content)
	}
}

func formatMeta(names, contents []string) []metaElem {
	max := stringsplus.GetMaxLength(names) + 1
	meta := make([]metaElem, len(names))
	for i, name := range names {
		if len(name) < max {
			name = name + ":" + strings.Repeat(" ", max-len(name))
			meta[i] = metaElem{name, contents[i]}
		}
	}
	return meta
}

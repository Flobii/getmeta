package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/flobii/getmeta/stringsplus"
	"github.com/gocolly/colly"
)

type metaelem struct {
	Name    string
	Content string
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println("Wrong number of arguments")
		return
	}
	url := os.Args[1]
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

func formatMeta(names, contents []string) []metaelem {
	max := stringsplus.GetMaxLength(names) + 1
	meta := make([]metaelem, len(names))
	for i, name := range names {
		if len(name) < max {
			name = name + ":" + strings.Repeat(" ", max-len(name))
			meta[i] = metaelem{name, contents[i]}
		}
	}
	return meta
}

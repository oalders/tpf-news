package main

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type BlogPost struct {
	Author  string
	Content string
	Date    time.Time
	Path    string
	Title   string
	Tags    []string
}

const doc = `
---
title: "{{.Title}}"
author: {{.Author}}
type: post
date: {{.Date}}
url: "{{.Path}}"
categories:
{{range .Tags}} - {{ . }}
{{end}}
---
{{ .Content }}
`

func main() {
	tmpl, err := template.New("test").Parse(doc)
	if err != nil {
		log.Fatal(err)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("news.perlfoundation.org"),
	)
	c.CacheDir = ".cache"
	c.AllowURLRevisit = false

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "/post") && !strings.HasPrefix(link, "/list") {
			return
		}

		_ = c.Visit(e.Request.AbsoluteURL(link))
	})

	dateRegex := regexp.MustCompile(`\d\d-\w\w\w-\d\d\d\d`)
	c.OnHTML(`div.blog-main`, func(e *colly.HTMLElement) {
		meta := e.DOM.Find("p.blog-post-meta")

		date := strings.TrimSpace(meta.Find("i").First().Text())
		dateMatch := dateRegex.Find([]byte(date))

		content := meta.Next()
		blogHTML, err := content.Html()
		if err != nil {
			log.Fatal(err)
		}
		tags := content.Next().Next().Next()
		var allTags []string
		tags.Find(`a[href]`).Each(func(i int, s *goquery.Selection) {
			allTags = append(allTags, s.Text())
		})

		post := BlogPost{
			Author:  meta.Find("a").First().Text(),
			Content: html.UnescapeString(blogHTML),
			Path:    e.Request.URL.Path,
			Tags:    allTags,
			Title:   strings.TrimSpace(e.DOM.Find("h2.blog-post-title").First().Text()),
		}

		post.Title = strings.ReplaceAll(post.Title, "\"", "\\\"")

		if len(dateMatch) > 0 {
			// 21-Jan-2020
			date, err := time.Parse("02-Jan-2006", string(dateMatch))
			if err == nil {
				post.Date = date
			} else {
				fmt.Printf("parse date: %v", err)
			}
		}
		fmt.Printf("%+v\n", post.Tags)

		f, err := os.Create(filepath.Join(
			"quickstart/content/post/", filepath.Base(post.Path)+".md"),
		)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)

		err = tmpl.Execute(w, post)
		if err != nil {
			log.Fatal(err)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	startUrl := "https://news.perlfoundation.org/list?limit=500&page=1"
	err = c.Visit(startUrl)
	if err != nil {
		log.Fatal(err)
	}
}

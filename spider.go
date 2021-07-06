package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"regexp"
	"strconv"
	"time"
)

const ShotPath = "shot"

// Mission the crawl mission
type Mission struct {
	Name    string
	Url     string
	Timeout uint8
	Next    []*Next
}

// Next the next config of mission
type Next struct {
	Rule     string
	RuleType string
	UrlPart  string
}

// Prey the prey of spider
type Prey struct {
	Html string
}

// Netting rod build grub mission
func Netting() *Mission {
	m := &Mission{
		Name:    "test",
		Url:     "https://www.wuyiting.cn/",
		Timeout: 0,
		Next:    nil,
	}

	var next []*Next
	next = append(next, &Next{
		Rule:     `/blog/([0-7])`,
		RuleType: "regex",
		UrlPart:  "https://www.wuyiting.cn/blog/",
	})

	m.Next = next
	return m
}

func Grub(m *Mission) {
	spider := rod.New().MustConnect()
	fmt.Println("get url: " + m.Url)
	page := spider.MustPage(m.Url)
	page.MustWaitLoad().MustScreenshot("a.png")
	html, _ := page.HTML()
	page.Close()
	if m.Next != nil && len(m.Next) > 0 {
		for _, next := range m.Next {
			urlList := Regex(next.Rule, html)
			for i, urlPart := range urlList {
				time.Sleep(time.Second * 2)
				url := "https://www.wuyiting.cn" + urlPart
				fmt.Println("get url: " + url)
				childPage := spider.MustPage(url)
				childPage.MustWaitLoad().MustScreenshot(strconv.Itoa(i) + ".png")
				childPage.Close()
			}
		}
	}
}

func Regex(rule string, html string) []string {
	pattern := regexp.MustCompile(rule)
	return pattern.FindAllString(html, -1)
}

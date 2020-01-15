package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// 有点消息队列的意思，chan最有用的功能就是阻塞！
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	<- tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	urls := os.Args[1:]
	worklist := make(chan []string)
	seen := make(map[string]bool)
	go func() { worklist <- urls }()
	for v := range worklist {
		for _, url := range v {
			if !seen[url] {
				seen[url] = true
				// 这里为什么不能这么写？？
				// 因为这样你的url就不知道拿的是哪个循环中的值
				// 可能会漏值，也可能是slice的最后一个值
				// go func() {worklist <- crawl(url)}()
				go func(url string) { worklist <- crawl(url) }(url)
			}
		}
	}
}

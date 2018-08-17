package rss

import (
	"testing"
	"fmt"
	"github.com/mmcdole/gofeed"
)

// rss 生成
// https://github.com/gorilla/feeds
func TestName(t *testing.T) {
	rssUrl := "http://feed.cnblogs.com/blog/u/222586/rss"
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rssUrl)
	fmt.Println(feed.Title)
	fmt.Println(feed.String())
}

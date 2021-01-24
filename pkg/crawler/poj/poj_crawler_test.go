package poj

import (
	"fmt"
	"testing"
)

func TestPOJCrawler_Crawl(t *testing.T) {
	crawler := &POJCrawler{}
	problem, err := crawler.Crawl("1001")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("problem: %v\n", problem.Hint)
}

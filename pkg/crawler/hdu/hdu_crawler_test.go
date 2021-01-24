package hdu

import (
	"fmt"
	"testing"
)

func TestHDUCrawler_Crawl(t *testing.T) {
	crawler := &HDUCrawler{}
	problem, err := crawler.Crawl("1001")
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("problem: %v", problem)
}

package hdu

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHDUCrawler_Crawl(t *testing.T) {
	crawler := &HDUCrawler{}
	problem, err := crawler.Crawl("1000")
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	str, _ := json.Marshal(problem)
	fmt.Printf(string(str))
}

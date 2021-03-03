package service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_crawler/pkg/bootstrap/database"
	"testing"
)

func TestCrawlerService_CrawlProblem(t *testing.T) {
	database.Init()
	service := &CrawlerService{}
	err := service.CrawlProblem(remote_oj.POJ, "1000")
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
}

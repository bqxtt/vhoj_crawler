package service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_crawler/pkg/bootstrap/database"
	"strconv"
	"testing"
)

func TestCrawlerService_CrawlProblem(t *testing.T) {
	database.Init()
	service := &CrawlerService{}
	for i := 1000; i < 1050; i++ {
		err := service.CrawlProblem(remote_oj.HDU, strconv.Itoa(i))
		if err != nil {
			fmt.Printf("%v err: %v\n", i, err)
		}
	}
	//err := service.CrawlProblem(remote_oj.HDU, "1000")
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//}
}

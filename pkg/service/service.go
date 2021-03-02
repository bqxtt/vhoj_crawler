package service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_crawler/pkg/holder"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/problem_mapper"
)

type CrawlerService struct {
}

func (CrawlerService) CrawlProblem(remoteOJ remote_oj.RemoteOJ, problemId string) error {
	crawler, exist := holder.Crawlers[remoteOJ]
	if exist == false {
		return fmt.Errorf("remote oj: %v, crawler not exist", remoteOJ)
	}
	rawProblem, err := crawler.Crawl(problemId)
	if err != nil {
		return err
	}
	_, err = problem_mapper.ProblemMapper.AddOrModifyRawProblem(rawProblem)
	if err != nil {
		return err
	}
	//todo group problem
	return nil
}

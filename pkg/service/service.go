package service

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_crawler/pkg/holder"
	"github.com/ecnuvj/vhoj_db/pkg/dao/mapper/problem_mapper"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
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
	if rawProblem.ID == 0 {
		fmt.Println("fuck")
	}

	//todo group problem
	_, err = problem_mapper.ProblemMapper.AddOrModifyProblemGroup(&model.ProblemGroup{
		RawProblemId:    rawProblem.ID,
		GroupId:         rawProblem.ID,
		MainProblem:     true,
		RemoteOJ:        remote_oj.HDU,
		RemoteProblemId: problemId,
	})
	if err != nil {
		fmt.Printf("problem group error: %v", err)
	}
	_, _ = problem_mapper.ProblemMapper.AddOrModifyProblem(&model.Problem{
		GroupId:      rawProblem.ID,
		RawProblemId: rawProblem.ID,
	})
	if err != nil {
		fmt.Printf("problem error: %v", err)
	}
	return nil
}

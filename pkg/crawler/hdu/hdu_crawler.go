package hdu

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_crawler/pkg/crawler"
	"github.com/ecnuvj/vhoj_crawler/pkg/utils/http_utils"
	"github.com/ecnuvj/vhoj_db/pkg/dao/model"
	"log"
)

var HduCrawler = &HDUCrawler{}

type HDUCrawler struct {
	crawler.DefaultCrawlerImpl
}

func (H *HDUCrawler) Crawl(problemId string) (*model.RawProblem, error) {
	err := H.PreValidate(problemId)
	if err != nil {
		return nil, err
	}
	rawContent, err := http_utils.Download("GET", H.GetProblemUrl(problemId), "gbk")
	if err != nil {
		return nil, err
	}
	//fmt.Println(rawContent)
	return H.ParseProblemInfo(rawContent, problemId)
}

func (H *HDUCrawler) GetHost() string {
	return remote_oj.HDUInfo.Host
}

func (H *HDUCrawler) GetProblemUrl(problemId string) string {
	return remote_oj.HDUInfo.Host + remote_oj.HDUInfo.ProblemUrl + problemId
}

func (H *HDUCrawler) PreValidate(problemId string) error {
	return nil
}

func (H *HDUCrawler) ParseProblemInfo(html string, problemId string) (*model.RawProblem, error) {
	rawProblem := &model.RawProblem{}
	var err error
	rawProblem.Title, err = http_utils.ParseHtmlReg(`color:#1A5CC8'>([\s\S]*?)</h1>`, html)
	if err != nil {
		return nil, err
	}
	timeLimit, err := http_utils.ParseHtmlReg(`(\d*) MS`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.TimeLimit = timeLimit

	memoryLimit, err := http_utils.ParseHtmlReg(`/(\d*) K`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.MemoryLimit = memoryLimit

	rawProblem.Description, err = http_utils.ParseHtmlReg(`>Problem Description</div>([\s\S]*?)<[^<>]*?panel_bottom[^<>]*?>`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.Input, err = http_utils.ParseHtmlReg(`>Input</div>([\s\S]*?)<[^<>]*?panel_bottom[^<>]*?>`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.Output, err = http_utils.ParseHtmlReg(`>Output</div>([\s\S]*?)<[^<>]*?panel_bottom[^<>]*?>`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.SampleInput, err = http_utils.ParseHtmlReg(`>Sample Input</div>([\s\S]*?)<[^<>]*?panel_bottom[^<>]*?>`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.SampleOutput, err = http_utils.ParseHtmlReg(`>Sample Output</div>([\s\S]*?)<[^<>]*?panel_bottom[^<>]*?>`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.Source, err = http_utils.ParseHtmlReg(`>Author</div>([\s\S]*?)<[^<>]*?panel_bottom[^<>]*?>`, html)
	if err != nil {
		log.Print("author no match")
		rawProblem.Source, _ = http_utils.ParseHtmlReg(`>Source</div>([\s\S]*?)<[^<>]*?panel_bottom[^<>]*?>`, html)
	}

	rawProblem.RemoteOJ = remote_oj.HDU
	rawProblem.RemoteProblemId = problemId

	return rawProblem, nil
}

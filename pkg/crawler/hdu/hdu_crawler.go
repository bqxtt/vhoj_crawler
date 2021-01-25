package hdu

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_crawler/pkg/common/entity"
	"github.com/bqxtt/vhoj_crawler/pkg/crawler"
	"github.com/bqxtt/vhoj_crawler/pkg/utils/http_utils"
	"strconv"
)

type HDUCrawler struct {
	crawler.DefaultCrawlerImpl
}

func (H *HDUCrawler) Crawl(problemId string) (*entity.RawProblem, error) {
	err := H.PreValidate(problemId)
	if err != nil {
		return nil, err
	}
	rawContent, err := http_utils.Download("GET", H.GetProblemUrl(problemId))
	if err != nil {
		return nil, err
	}
	//fmt.Println(rawContent)
	return H.ParseProblemInfo(rawContent, problemId)
}

func (H *HDUCrawler) GetHost() string {
	return constants.HDUInfo.Host
}

func (H *HDUCrawler) GetProblemUrl(problemId string) string {
	return constants.HDUInfo.Host + constants.HDUInfo.ProblemUrl + problemId
}

func (H *HDUCrawler) PreValidate(problemId string) error {
	return nil
}

func (H *HDUCrawler) ParseProblemInfo(html string, problemId string) (*entity.RawProblem, error) {
	rawProblem := &entity.RawProblem{}
	var err error
	rawProblem.Title, err = http_utils.ParseHtmlReg(`color:#1A5CC8'>([\s\S]*?)</h1>`, html)
	if err != nil {
		return nil, err
	}
	timeLimit, err := http_utils.ParseHtmlReg(`(\d*) MS`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.TimeLimit, err = strconv.ParseInt(timeLimit, 10, 64)
	if err != nil {
		return nil, err
	}
	memoryLimit, err := http_utils.ParseHtmlReg(`/(\d*) K`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.MemoryLimit, err = strconv.ParseInt(memoryLimit, 10, 64)
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}
	return rawProblem, nil
}

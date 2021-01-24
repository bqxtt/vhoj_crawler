package poj

import (
	"github.com/vhoj_crawler/pkg/common/constant/poj"
	"github.com/vhoj_crawler/pkg/common/entity"
	"github.com/vhoj_crawler/pkg/crawler"
	"github.com/vhoj_crawler/pkg/utils/http_utils"
	"log"
	"strconv"
)

type POJCrawler struct {
	crawler.DefaultCrawlerImpl
}

func (P *POJCrawler) Crawl(problemId string) (*entity.RawProblem, error) {
	err := P.PreValidate(problemId)
	if err != nil {
		return nil, err
	}
	rawContent, err := http_utils.Download("GET", P.GetProblemUrl(problemId))
	if err != nil {
		return nil, err
	}
	//fmt.Println(rawContent)
	return P.ParseProblemInfo(rawContent, problemId)
}

func (P *POJCrawler) GetHost() string {
	return poj.HOST
}

func (P *POJCrawler) GetProblemUrl(problemId string) string {
	return poj.HOST + poj.PROBLEM_URL + problemId
}

func (P *POJCrawler) PreValidate(problemId string) error {
	return nil
}

func (P *POJCrawler) ParseProblemInfo(html string, problemId string) (*entity.RawProblem, error) {
	rawProblem := &entity.RawProblem{}
	var err error
	rawProblem.Title, err = http_utils.ParseHtmlReg(`<title>\d{3,} -- ([\s\S]*?)</title>`, html)
	if err != nil {
		return nil, err
	}
	timeLimit, err := http_utils.ParseHtmlReg(`<b>Time Limit:</b> (\d{3,})MS</td>`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.TimeLimit, err = strconv.ParseInt(timeLimit, 10, 64)
	if err != nil {
		return nil, err
	}
	memoryLimit, err := http_utils.ParseHtmlReg(`<b>Memory Limit:</b> (\d{2,})K</td>`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.MemoryLimit, err = strconv.ParseInt(memoryLimit, 10, 64)
	if err != nil {
		return nil, err
	}
	rawProblem.Description, err = http_utils.ParseHtmlReg(`<p class="pst">Description</p><[\s\S]*?>([\s\S]*?)<[\s\S]*?><p class="pst">`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.Input, err = http_utils.ParseHtmlReg(`<p class="pst">Input</p><[\s\S]*?>([\s\S]*?)<[\s\S]*?><p class="pst">`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.Output, err = http_utils.ParseHtmlReg(`<p class="pst">Output</p><[\s\S]*?>([\s\S]*?)<[\s\S]*?><p class="pst">`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.SampleInput, err = http_utils.ParseHtmlReg(`<p class="pst">Sample Input</p><[\s\S]*?>([\s\S]*?)<[\s\S]*?><p class="pst">`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.SampleOutput, err = http_utils.ParseHtmlReg(`<p class="pst">Sample Output</p><[\s\S]*?>([\s\S]*?)<[\s\S]*?><p class="pst">`, html)
	if err != nil {
		return nil, err
	}
	rawProblem.Hint, err = http_utils.ParseHtmlReg(`<p class="pst">Hint</p>([\s\S]*?)<p class="pst">`, html)
	if err != nil {
		if err.Error() == "no match" {
			log.Print("hint no match")
		} else {
			return nil, err
		}
	}
	rawProblem.Source, err = http_utils.ParseHtmlReg(`<p class="pst">Source</p>([\s\S]*?)</td></tr></tbody></table>`, html)
	if err != nil {
		if err.Error() == "no match" {
			log.Print("source no match")
		} else {
			return nil, err
		}
	}

	//fmt.Printf("%v", rawProblem)
	return rawProblem, nil
}

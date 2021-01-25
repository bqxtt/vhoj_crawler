package crawler

import "github.com/bqxtt/vhoj_common/pkg/common/entity"

type ICrawler interface {
	Crawl(problemId string) (*entity.RawProblem, error)
	GetHost() string
	GetProblemUrl(problemId string) string
	PreValidate(problemId string) error
	ParseProblemInfo(html string, problemId string) (*entity.RawProblem, error)
	mustEmbedDefaultCrawler()
}

type DefaultCrawlerImpl struct{}

func (DefaultCrawlerImpl) Crawl(problemId string) (*entity.RawProblem, error) {
	panic("implement me")
}

func (DefaultCrawlerImpl) GetHost() string {
	panic("implement me")
}

func (DefaultCrawlerImpl) GetProblemUrl(problemId string) string {
	panic("implement me")
}

func (DefaultCrawlerImpl) PreValidate(problemId string) error {
	panic("implement me")
}
func (DefaultCrawlerImpl) ParseProblemInfo(html string, problemId string) (*entity.RawProblem, error) {
	panic("implement me")
}

func (DefaultCrawlerImpl) mustEmbedDefaultCrawler() {

}

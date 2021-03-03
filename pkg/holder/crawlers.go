package holder

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_crawler/pkg/crawler"
	"github.com/ecnuvj/vhoj_crawler/pkg/crawler/hdu"
	"github.com/ecnuvj/vhoj_crawler/pkg/crawler/poj"
)

var Crawlers map[remote_oj.RemoteOJ]crawler.ICrawler

func init() {
	Crawlers = make(map[remote_oj.RemoteOJ]crawler.ICrawler, 0)
	Crawlers[remote_oj.POJ] = poj.PojCrawler
	Crawlers[remote_oj.HDU] = hdu.HduCrawler
}
